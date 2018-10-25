module Resource
    exposing
        ( Flags
        , Msg(..)
        , Model
        , init
        , changeToResource
        , update
        , updateWithMessage
        , view
        , viewCheckbox
        , viewPin
        , viewVersionHeader
        , viewVersionBody
        , subscriptions
        , PauseChangingOrErrored(..)
        , PinState(..)
        )

import Concourse
import Concourse.BuildStatus
import Concourse.Pagination exposing (Pagination, Paginated, Page, equal)
import Concourse.Resource
import Css
import Dict
import DictView
import Date exposing (Date)
import Date.Format
import Duration exposing (Duration)
import Erl
import Html.Styled as Html exposing (Html)
import Html.Styled.Attributes exposing (class, css, href, id, title)
import Html.Styled.Events exposing (onClick, onMouseEnter, onMouseLeave, onMouseOver, onMouseOut)
import Http
import Maybe.Extra as ME
import Navigation
import NewTopBar.Styles as Styles
import StrictEvents
import Task exposing (Task)
import Time exposing (Time)
import LoginRedirect
import RemoteData exposing (WebData)
import UpdateMsg exposing (UpdateMsg)


type alias Ports =
    { title : String -> Cmd Msg
    }


type alias Model =
    { ports : Ports
    , now : Maybe Time.Time
    , resourceIdentifier : Concourse.ResourceIdentifier
    , resource : WebData Concourse.Resource
    , pausedChanging : PauseChangingOrErrored
    , versionedResources : Paginated Concourse.VersionedResource
    , currentPage : Maybe Page
    , versionedUIStates : Dict.Dict Int VersionUIState
    , csrfToken : String
    , showPinBarTooltip : Bool
    }


type alias VersionUIState =
    { changingErrored : Bool
    , expanded : Bool
    , inputTo : List Concourse.Build
    , outputOf : List Concourse.Build
    , showTooltip : Bool
    }


type PauseChangingOrErrored
    = Stable
    | Changing
    | Errored


type PinState
    = Unpinned
    | Pinned
    | Disabled


type Msg
    = Noop
    | AutoupdateTimerTicked Time
    | ResourceFetched (Result Http.Error Concourse.Resource)
    | TogglePaused
    | PausedToggled (Result Http.Error ())
    | VersionedResourcesFetched (Maybe Page) (Result Http.Error (Paginated Concourse.VersionedResource))
    | LoadPage Page
    | ClockTick Time.Time
    | ToggleVersionedResource Int
    | VersionedResourceToggled Int (Result Http.Error ())
    | ExpandVersionedResource Int
    | InputToFetched Int (Result Http.Error (List Concourse.Build))
    | OutputOfFetched Int (Result Http.Error (List Concourse.Build))
    | NavTo String
    | TogglePinBarTooltip
    | ToggleVersionTooltip Int


type alias Flags =
    { teamName : String
    , pipelineName : String
    , resourceName : String
    , paging : Maybe Concourse.Pagination.Page
    , csrfToken : String
    }


init : Ports -> Flags -> ( Model, Cmd Msg )
init ports flags =
    let
        ( model, cmd ) =
            changeToResource flags
                { resourceIdentifier =
                    { teamName = flags.teamName
                    , pipelineName = flags.pipelineName
                    , resourceName = flags.resourceName
                    }
                , resource = RemoteData.NotAsked
                , pausedChanging = Stable
                , currentPage = Nothing
                , versionedResources =
                    { content = []
                    , pagination =
                        { previousPage = Nothing
                        , nextPage = Nothing
                        }
                    }
                , versionedUIStates = Dict.empty
                , ports = ports
                , now = Nothing
                , csrfToken = flags.csrfToken
                , showPinBarTooltip = False
                }
    in
        ( model
        , Cmd.batch
            [ fetchResource model.resourceIdentifier
            , cmd
            ]
        )


changeToResource : Flags -> Model -> ( Model, Cmd Msg )
changeToResource flags model =
    ( { model
        | currentPage = flags.paging
        , versionedResources =
            { content = []
            , pagination =
                { previousPage = Nothing
                , nextPage = Nothing
                }
            }
      }
    , fetchVersionedResources model.resourceIdentifier flags.paging
    )


updateWithMessage : Msg -> Model -> ( Model, Cmd Msg, Maybe UpdateMsg )
updateWithMessage message model =
    let
        ( mdl, msg ) =
            update message model
    in
        case mdl.resource of
            RemoteData.Failure _ ->
                ( mdl, msg, Just UpdateMsg.NotFound )

            _ ->
                ( mdl, msg, Nothing )


update : Msg -> Model -> ( Model, Cmd Msg )
update action model =
    case action of
        Noop ->
            ( model, Cmd.none )

        AutoupdateTimerTicked timestamp ->
            ( model
            , Cmd.batch <|
                List.append
                    [ fetchResource model.resourceIdentifier
                    , fetchVersionedResources model.resourceIdentifier model.currentPage
                    ]
                <|
                    updateExpandedProperties model
            )

        ResourceFetched (Ok resource) ->
            ( { model | resource = RemoteData.Success resource }
            , model.ports.title <| resource.name ++ " - "
            )

        ResourceFetched (Err err) ->
            case Debug.log ("failed to fetch resource") (err) of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, LoginRedirect.requestLoginRedirect "" )
                    else if status.code == 404 then
                        ( { model | resource = RemoteData.Failure err }, Cmd.none )
                    else
                        ( model, Cmd.none )

                _ ->
                    ( model, Cmd.none )

        TogglePaused ->
            case model.resource |> RemoteData.toMaybe of
                Nothing ->
                    ( model, Cmd.none )

                Just r ->
                    ( { model
                        | pausedChanging = Changing
                        , resource = RemoteData.Success { r | paused = not r.paused }
                      }
                    , if r.paused then
                        unpauseResource model.resourceIdentifier model.csrfToken
                      else
                        pauseResource model.resourceIdentifier model.csrfToken
                    )

        PausedToggled (Ok ()) ->
            ( { model | pausedChanging = Stable }, Cmd.none )

        PausedToggled (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, LoginRedirect.requestLoginRedirect "" )
                    else
                        ( model, Cmd.none )

                _ ->
                    case model.resource |> RemoteData.toMaybe of
                        Nothing ->
                            ( model, Cmd.none )

                        Just r ->
                            ( { model
                                | pausedChanging = Errored
                                , resource = RemoteData.Success { r | paused = not r.paused }
                              }
                            , Cmd.none
                            )

        VersionedResourcesFetched requestedPage (Ok paginated) ->
            let
                fetchedPage =
                    permalink paginated.content

                newModel =
                    \newPage ->
                        { model
                            | versionedResources = paginated
                            , currentPage = newPage
                        }

                chosenModelWith =
                    \requestedPageUnwrapped ->
                        case model.currentPage of
                            Nothing ->
                                newModel <| Just fetchedPage

                            Just page ->
                                if Concourse.Pagination.equal page requestedPageUnwrapped then
                                    newModel <| requestedPage
                                else
                                    model
            in
                case requestedPage of
                    Nothing ->
                        ( newModel (Just fetchedPage), Cmd.none )

                    Just requestedPageUnwrapped ->
                        ( chosenModelWith requestedPageUnwrapped
                        , Cmd.none
                        )

        VersionedResourcesFetched _ (Err err) ->
            flip always (Debug.log ("failed to fetch versioned resources") (err)) <|
                ( model, Cmd.none )

        LoadPage page ->
            ( { model
                | currentPage = Just page
              }
            , Cmd.batch
                [ fetchVersionedResources model.resourceIdentifier <| Just page
                , Navigation.newUrl <| paginationRoute model.resourceIdentifier page
                ]
            )

        ToggleVersionedResource versionID ->
            let
                versionedResourceIdentifier =
                    { teamName = model.resourceIdentifier.teamName
                    , pipelineName = model.resourceIdentifier.pipelineName
                    , resourceName = model.resourceIdentifier.resourceName
                    , versionID = versionID
                    }

                versionedResource =
                    List.head <|
                        List.filter (checkForVersionID versionID) model.versionedResources.content
            in
                ( model
                , case versionedResource of
                    Just vr ->
                        if vr.enabled then
                            disableVersionedResource versionedResourceIdentifier model.csrfToken
                        else
                            enableVersionedResource versionedResourceIdentifier model.csrfToken

                    Nothing ->
                        Cmd.none
                )

        VersionedResourceToggled versionID (Ok ()) ->
            let
                oldState =
                    getState versionID model.versionedUIStates

                newState =
                    { oldState
                        | changingErrored = False
                    }

                oldVRs =
                    model.versionedResources

                oldContent =
                    model.versionedResources.content
            in
                ( { model
                    | versionedResources =
                        { oldVRs
                            | content = updateMatchingMember versionID oldContent
                        }
                    , versionedUIStates = setState versionID newState model.versionedUIStates
                  }
                , Cmd.none
                )

        VersionedResourceToggled versionID (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, LoginRedirect.requestLoginRedirect "" )
                    else
                        ( model, Cmd.none )

                _ ->
                    let
                        oldState =
                            getState versionID model.versionedUIStates

                        newState =
                            { oldState
                                | expanded = not oldState.expanded
                                , changingErrored = True
                            }
                    in
                        ( { model
                            | versionedUIStates = setState versionID newState model.versionedUIStates
                          }
                        , Cmd.none
                        )

        ExpandVersionedResource versionID ->
            let
                versionedResourceIdentifier =
                    { teamName = model.resourceIdentifier.teamName
                    , pipelineName = model.resourceIdentifier.pipelineName
                    , resourceName = model.resourceIdentifier.resourceName
                    , versionID = versionID
                    }

                oldState =
                    getState versionID model.versionedUIStates

                newState =
                    { oldState
                        | expanded = not oldState.expanded
                    }
            in
                ( { model
                    | versionedUIStates = setState versionID newState model.versionedUIStates
                  }
                , if newState.expanded then
                    Cmd.batch
                        [ fetchInputTo versionedResourceIdentifier
                        , fetchOutputOf versionedResourceIdentifier
                        ]
                  else
                    Cmd.none
                )

        InputToFetched _ (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, LoginRedirect.requestLoginRedirect "" )
                    else
                        ( model, Cmd.none )

                _ ->
                    ( model, Cmd.none )

        InputToFetched versionID (Ok builds) ->
            let
                oldState =
                    getState versionID model.versionedUIStates

                newState =
                    { oldState
                        | inputTo = builds
                    }
            in
                ( { model
                    | versionedUIStates = setState versionID newState model.versionedUIStates
                  }
                , Cmd.none
                )

        OutputOfFetched _ (Err err) ->
            case err of
                Http.BadStatus { status } ->
                    if status.code == 401 then
                        ( model, LoginRedirect.requestLoginRedirect "" )
                    else
                        ( model, Cmd.none )

                _ ->
                    ( model, Cmd.none )

        ClockTick now ->
            ( { model | now = Just now }, Cmd.none )

        OutputOfFetched versionID (Ok builds) ->
            let
                oldState =
                    getState versionID model.versionedUIStates

                newState =
                    { oldState
                        | outputOf = builds
                    }
            in
                ( { model
                    | versionedUIStates = setState versionID newState model.versionedUIStates
                  }
                , Cmd.none
                )

        NavTo url ->
            ( model, Navigation.newUrl url )

        TogglePinBarTooltip ->
            ( { model
                | showPinBarTooltip =
                    if
                        model.resource
                            |> RemoteData.map .pinnedInConfig
                            |> RemoteData.withDefault False
                    then
                        not model.showPinBarTooltip
                    else
                        False
              }
            , Cmd.none
            )

        ToggleVersionTooltip versionID ->
            let
                oldState =
                    getState versionID model.versionedUIStates

                newState =
                    { oldState
                        | showTooltip = not oldState.showTooltip
                    }
            in
                ( { model
                    | versionedUIStates = setState versionID newState model.versionedUIStates
                  }
                , Cmd.none
                )


permalink : List Concourse.VersionedResource -> Page
permalink versionedResources =
    case List.head versionedResources of
        Nothing ->
            { direction = Concourse.Pagination.Since 0
            , limit = 100
            }

        Just version ->
            { direction = Concourse.Pagination.From version.id
            , limit = List.length versionedResources
            }


paginationRoute : Concourse.ResourceIdentifier -> Page -> String
paginationRoute rid page =
    let
        ( param, boundary ) =
            case page.direction of
                Concourse.Pagination.Since bound ->
                    ( "since", Basics.toString bound )

                Concourse.Pagination.Until bound ->
                    ( "until", Basics.toString bound )

                Concourse.Pagination.From bound ->
                    ( "from", Basics.toString bound )

                Concourse.Pagination.To bound ->
                    ( "to", Basics.toString bound )

        parsedRoute =
            Erl.parse <|
                "/teams/"
                    ++ rid.teamName
                    ++ "/pipelines/"
                    ++ rid.pipelineName
                    ++ "/resources/"
                    ++ rid.resourceName

        newParsedRoute =
            Erl.addQuery param boundary <| Erl.addQuery "limit" (Basics.toString page.limit) parsedRoute
    in
        Erl.toString newParsedRoute


view : Model -> Html Msg
view model =
    case model.resource |> RemoteData.toMaybe of
        Just resource ->
            let
                ( checkStatus, checkMessage, stepBody ) =
                    if resource.failingToCheck then
                        if not (String.isEmpty resource.checkSetupError) then
                            ( "fr errored fa fa-fw fa-exclamation-triangle"
                            , "checking failed"
                            , [ Html.div [ class "step-body" ]
                                    [ Html.pre [] [ Html.text resource.checkSetupError ]
                                    ]
                              ]
                            )
                        else
                            ( "fr errored fa fa-fw fa-exclamation-triangle"
                            , "checking failed"
                            , [ Html.div [ class "step-body" ]
                                    [ Html.pre [] [ Html.text resource.checkError ]
                                    ]
                              ]
                            )
                    else
                        ( "fr succeeded fa fa-fw fa-check", "checking successfully", [] )

                ( paused, pausedIcon, aria, onClickEvent ) =
                    case ( resource.paused, model.pausedChanging ) of
                        ( _, Changing ) ->
                            ( "loading", "fa-spin fa-circle-o-notch", "", Noop )

                        ( True, Errored ) ->
                            ( "errored", "fa-play", "Resume Resource Checking", TogglePaused )

                        ( False, Errored ) ->
                            ( "errored", "fa-pause", "Pause Resource Checking", TogglePaused )

                        ( True, Stable ) ->
                            ( "enabled", "fa-play", "Resume Resource Checking", TogglePaused )

                        ( False, Stable ) ->
                            ( "disabled", "fa-pause", "Pause Resource Checking", TogglePaused )

                ( previousButtonClass, previousButtonEvent ) =
                    case model.versionedResources.pagination.previousPage of
                        Nothing ->
                            ( "btn-page-link prev disabled", Noop )

                        Just pp ->
                            ( "btn-page-link prev", LoadPage pp )

                ( nextButtonClass, nextButtonEvent ) =
                    case model.versionedResources.pagination.nextPage of
                        Nothing ->
                            ( "btn-page-link next disabled", Noop )

                        Just np ->
                            let
                                updatedPage =
                                    { np
                                        | limit = 100
                                    }
                            in
                                ( "btn-page-link next", LoadPage updatedPage )

                lastCheckedView =
                    case ( model.now, resource.lastChecked ) of
                        ( Just now, Just date ) ->
                            viewLastChecked now date

                        ( _, _ ) ->
                            Html.text ""

                pinBar =
                    Html.div
                        ([ css
                            [ Css.flexGrow (Css.num 1)
                            , Css.border3 (Css.px 1)
                                Css.solid
                                (if ME.isJust resource.pinnedVersion then
                                    Css.hex "03dac4"
                                 else
                                    Css.hex "3d3c3c"
                                )
                            , Css.margin (Css.px 10)
                            , Css.paddingLeft (Css.px 7)
                            , Css.displayFlex
                            , Css.alignItems Css.center
                            , Css.position Css.relative
                            ]
                         , id "pin-bar"
                         ]
                            ++ (if resource.pinnedInConfig then
                                    [ onMouseEnter TogglePinBarTooltip
                                    , onMouseLeave TogglePinBarTooltip
                                    ]
                                else
                                    []
                               )
                        )
                        ([ Html.div
                            [ css
                                [ Css.backgroundImage
                                    (if ME.isJust resource.pinnedVersion then
                                        Css.url "/public/images/pin_ic_teal.svg"
                                     else
                                        Css.url "/public/images/pin_ic_grey.svg"
                                    )
                                , Css.backgroundRepeat Css.noRepeat
                                , Css.backgroundPosition2 (Css.pct 50) (Css.pct 50)
                                , Css.height (Css.px 15)
                                , Css.width (Css.px 15)
                                , Css.marginRight (Css.px 10)
                                ]
                            ]
                            []
                         , resource.pinnedVersion
                            |> Maybe.map viewVersion
                            |> Maybe.withDefault (Html.text "")
                         ]
                            ++ (if model.showPinBarTooltip then
                                    [ Html.div
                                        [ css
                                            [ Css.position Css.absolute
                                            , Css.top <| Css.px -10
                                            , Css.left <| Css.px 30
                                            , Css.backgroundColor <| Css.hex "9b9b9b"
                                            , Css.zIndex <| Css.int 2
                                            , Css.padding <| Css.px 5
                                            ]
                                        ]
                                        [ Html.text "pinned in pipeline config" ]
                                    ]
                                else
                                    []
                               )
                        )

                headerHeight =
                    60
            in
                Html.div []
                    [ Html.div
                        [ css
                            [ Css.height <| Css.px headerHeight
                            , Css.position Css.fixed
                            , Css.top <| Css.px Styles.pageHeaderHeight
                            , Css.displayFlex
                            , Css.alignItems Css.stretch
                            , Css.width <| Css.pct 100
                            , Css.zIndex <| Css.int 1
                            , Css.backgroundColor <| Css.hex "2a2929"
                            ]
                        ]
                        [ Html.h1
                            [ css
                                [ Css.fontWeight <| Css.int 700
                                , Css.marginLeft <| Css.px 18
                                , Css.displayFlex
                                , Css.alignItems Css.center
                                , Css.justifyContent Css.center
                                ]
                            ]
                            [ Html.text resource.name ]
                        , Html.div
                            [ css
                                [ Css.displayFlex
                                , Css.alignItems Css.center
                                , Css.justifyContent Css.center
                                , Css.marginLeft (Css.px 24)
                                ]
                            ]
                            [ lastCheckedView ]
                        , pinBar
                        , Html.div
                            [ class previousButtonClass
                            , onClick previousButtonEvent
                            , css [ Css.displayFlex, Css.alignItems Css.center ]
                            ]
                            [ Html.a [ class "arrow" ]
                                [ Html.i [ class "fa fa-arrow-left" ] []
                                ]
                            ]
                        , Html.div
                            [ class nextButtonClass
                            , onClick nextButtonEvent
                            , css [ Css.displayFlex, Css.alignItems Css.center ]
                            ]
                            [ Html.a [ class "arrow" ]
                                [ Html.i [ class "fa fa-arrow-right" ] []
                                ]
                            ]
                        ]
                    , Html.div
                        [ css
                            [ Css.padding3 (Css.px <| headerHeight + 10) (Css.px 10) (Css.px 10)
                            ]
                        ]
                        [ Html.div [ class "resource-check-status" ]
                            [ Html.div [ class "build-step" ]
                                (List.append
                                    [ Html.div [ class "header" ]
                                        [ Html.button
                                            [ class <| "btn-pause fl " ++ paused
                                            , Html.Styled.Attributes.attribute "aria-label" aria
                                            , title aria
                                            , onClick onClickEvent
                                            ]
                                            [ Html.i [ class <| "fa fa-fw " ++ pausedIcon ] []
                                            ]
                                        , Html.h3 [] [ Html.text checkMessage ]
                                        , Html.i [ class <| checkStatus ] []
                                        ]
                                    ]
                                    stepBody
                                )
                            ]
                        , viewVersionedResources
                            { versionedResources = model.versionedResources
                            , versionedUIStates = model.versionedUIStates
                            , isResourcePinnedInConfig = resource.pinnedInConfig
                            , pinnedVersion = resource.pinnedVersion
                            }
                        ]
                    ]

        Nothing ->
            Html.div [] []


checkForVersionID : Int -> Concourse.VersionedResource -> Bool
checkForVersionID versionID versionedResource =
    versionID == versionedResource.id


updateMatchingMember : Int -> List Concourse.VersionedResource -> List Concourse.VersionedResource
updateMatchingMember versionID versionedResources =
    List.map (switchEnabled versionID) versionedResources


switchEnabled : Int -> Concourse.VersionedResource -> Concourse.VersionedResource
switchEnabled versionID versionedResource =
    let
        wasEnabled =
            versionedResource.enabled
    in
        if versionID == versionedResource.id then
            { versionedResource
                | enabled = not wasEnabled
            }
        else
            versionedResource


viewVersionedResources :
    { a
        | versionedResources : Paginated Concourse.VersionedResource
        , versionedUIStates : Dict.Dict Int VersionUIState
        , isResourcePinnedInConfig : Bool
        , pinnedVersion : Maybe Concourse.Version
    }
    -> Html Msg
viewVersionedResources { versionedResources, isResourcePinnedInConfig, versionedUIStates, pinnedVersion } =
    versionedResources.content
        |> List.map
            (\vr ->
                viewVersionedResource
                    { versionedResource = vr
                    , pinState =
                        (if isResourcePinnedInConfig then
                            Disabled
                         else if pinnedVersion == Just vr.version then
                            Pinned
                         else if pinnedVersion == Nothing then
                            Unpinned
                         else
                            Disabled
                        )
                    , state = getState vr.id versionedUIStates
                    }
            )
        |> Html.ul [ class "list list-collapsable list-enableDisable resource-versions" ]


viewVersionedResource :
    { versionedResource : Concourse.VersionedResource
    , state : VersionUIState
    , pinState : PinState
    }
    -> Html Msg
viewVersionedResource { versionedResource, pinState, state } =
    Html.li []
        ([ Html.div
            [ css
                [ Css.displayFlex
                , Css.margin2 (Css.px 5) Css.zero
                ]
            ]
            [ viewCheckbox versionedResource
            , viewPin
                { id = versionedResource.id
                , pinState = pinState
                , showTooltip = state.showTooltip
                }
            , viewVersionHeader versionedResource
            ]
         ]
            ++ (if state.expanded then
                    [ viewVersionBody
                        { inputTo = state.inputTo
                        , outputOf = state.outputOf
                        , metadata = versionedResource.metadata
                        , enabled = versionedResource.enabled
                        }
                    ]
                else
                    []
               )
        )


viewVersionBody :
    { a
        | inputTo : List Concourse.Build
        , outputOf : List Concourse.Build
        , metadata : Concourse.Metadata
        , enabled : Bool
    }
    -> Html Msg
viewVersionBody { inputTo, outputOf, metadata, enabled } =
    Html.div
        [ css
            ([ Css.displayFlex
             , Css.padding2 (Css.px 5) (Css.px 10)
             ]
                ++ (if enabled then
                        []
                    else
                        [ Css.opacity (Css.num 0.5) ]
                   )
            )
        ]
        [ Html.div [ class "vri" ] <|
            List.concat
                [ [ Html.div [ css [ Css.lineHeight <| Css.px 25 ] ] [ Html.text "inputs to" ] ]
                , viewBuilds <| listToMap inputTo
                ]
        , Html.div [ class "vri" ] <|
            List.concat
                [ [ Html.div [ css [ Css.lineHeight <| Css.px 25 ] ] [ Html.text "outputs of" ] ]
                , viewBuilds <| listToMap outputOf
                ]
        , Html.div [ class "vri metadata-container" ]
            [ Html.div [ class "list-collapsable-title" ] [ Html.text "metadata" ]
            , viewMetadata metadata
            ]
        ]


viewCheckbox : { a | enabled : Bool, id : Int } -> Html Msg
viewCheckbox { enabled, id } =
    Html.a
        [ Html.Styled.Attributes.attribute "aria-label" "Toggle Resource Version"
        , css
            ((if enabled then
                [ Css.backgroundImage (Css.url "/public/images/checkmark_ic.svg")
                , Css.backgroundColor (Css.hex "2ecc71")
                ]
              else
                [ Css.backgroundImage (Css.url "/public/images/x_ic.svg")
                , Css.backgroundColor (Css.hex "1e1d1d")
                ]
             )
                ++ [ Css.backgroundRepeat Css.noRepeat
                   , Css.backgroundPosition2 (Css.pct 50) (Css.pct 50)
                   , Css.marginRight (Css.px 5)
                   , Css.width (Css.px 25)
                   , Css.height (Css.px 25)
                   , Css.float Css.left
                   ]
            )
        , onClick <| ToggleVersionedResource id
        ]
        []


viewPin : { id : Int, pinState : PinState, showTooltip : Bool } -> Html Msg
viewPin { id, pinState, showTooltip } =
    Html.div
        [ Html.Styled.Attributes.attribute "aria-label" "Pin Resource Version"
        , css
            [ Css.position Css.relative
            , Css.backgroundImage
                (Css.url "/public/images/pin_ic_grey.svg")
            , Css.backgroundRepeat Css.noRepeat
            , Css.backgroundPosition2 (Css.pct 50) (Css.pct 50)
            , Css.marginRight (Css.px 5)
            , Css.width (Css.px 25)
            , Css.height (Css.px 25)
            , Css.float Css.left
            , Css.cursor Css.default
            , Css.backgroundColor
                (case pinState of
                    Pinned ->
                        Css.hex "#03dac4"

                    Disabled ->
                        Css.hex "1e1d1d80"

                    Unpinned ->
                        Css.hex "1e1d1d"
                )
            ]
        , onMouseOut <| ToggleVersionTooltip id
        , onMouseOver <| ToggleVersionTooltip id
        ]
        (if showTooltip then
            [ Html.div
                [ css
                    [ Css.position Css.absolute
                    , Css.bottom <| Css.px 25
                    , Css.backgroundColor <| Css.hex "9b9b9b"
                    , Css.zIndex <| Css.int 2
                    , Css.padding <| Css.px 5
                    , Css.width <| Css.px 170
                    ]
                ]
                [ Html.text "enable via pipeline config" ]
            ]
         else
            []
        )


viewVersionHeader : { a | id : Int, enabled : Bool, version : Concourse.Version } -> Html Msg
viewVersionHeader { id, enabled, version } =
    Html.div
        [ css
            [ Css.flexGrow <| Css.num 1
            , Css.backgroundColor <| Css.hex "1e1d1d"
            , Css.cursor Css.pointer
            , Css.displayFlex
            , Css.alignItems Css.center
            , Css.paddingLeft <| Css.px 10
            , Css.color <|
                Css.hex <|
                    if enabled then
                        "e6e7e8"
                    else
                        "e6e7e880"
            ]
        , onClick <| ExpandVersionedResource id
        ]
        [ viewVersion version ]


getState : Int -> Dict.Dict Int VersionUIState -> VersionUIState
getState versionID states =
    let
        resourceState =
            Dict.get versionID states
    in
        case resourceState of
            Nothing ->
                { changingErrored = False
                , expanded = False
                , inputTo = []
                , outputOf = []
                , showTooltip = False
                }

            Just rs ->
                rs


setState : Int -> VersionUIState -> Dict.Dict Int VersionUIState -> Dict.Dict Int VersionUIState
setState versionID newState states =
    Dict.insert versionID newState states


viewVersion : Concourse.Version -> Html Msg
viewVersion version =
    (Html.fromUnstyled << DictView.view)
        << Dict.map (\_ s -> Html.toUnstyled (Html.text s))
    <|
        version


viewMetadata : Concourse.Metadata -> Html Msg
viewMetadata metadata =
    Html.dl [ class "build-metadata" ]
        (List.concatMap viewMetadataField metadata)


viewMetadataField : Concourse.MetadataField -> List (Html a)
viewMetadataField field =
    [ Html.dt [] [ Html.text field.name ]
    , Html.dd []
        [ Html.pre [ class "metadata-field" ] [ Html.text field.value ]
        ]
    ]


listToMap : List Concourse.Build -> Dict.Dict String (List Concourse.Build)
listToMap builds =
    let
        insertBuild =
            \build dict ->
                let
                    jobName =
                        case build.job of
                            Nothing ->
                                Debug.crash "Jobless builds shouldn't appear on this page!" ""

                            Just job ->
                                job.jobName

                    oldList =
                        Dict.get jobName dict

                    newList =
                        case oldList of
                            Nothing ->
                                [ build ]

                            Just list ->
                                list ++ [ build ]
                in
                    Dict.insert jobName newList dict
    in
        List.foldr insertBuild Dict.empty builds


viewBuilds : Dict.Dict String (List Concourse.Build) -> List (Html Msg)
viewBuilds buildDict =
    List.concatMap (viewBuildsByJob buildDict) <| Dict.keys buildDict


viewLastChecked : Time -> Date -> Html a
viewLastChecked now date =
    let
        ago =
            Duration.between (Date.toTime date) now
    in
        Html.table []
            [ Html.tr
                []
                [ Html.td [] [ Html.text "checked" ]
                , Html.td [ title (Date.Format.format "%b %d %Y %I:%M:%S %p" date) ]
                    [ Html.span [] [ Html.text (Duration.format ago ++ " ago") ] ]
                ]
            ]


viewBuildsByJob : Dict.Dict String (List Concourse.Build) -> String -> List (Html Msg)
viewBuildsByJob buildDict jobName =
    let
        oneBuildToLi =
            \build ->
                let
                    link =
                        case build.job of
                            Nothing ->
                                ""

                            Just job ->
                                "/teams/" ++ job.teamName ++ "/pipelines/" ++ job.pipelineName ++ "/jobs/" ++ job.jobName ++ "/builds/" ++ build.name
                in
                    Html.li [ class <| Concourse.BuildStatus.show build.status ]
                        [ Html.a
                            [ Html.Styled.Attributes.fromUnstyled <| StrictEvents.onLeftClick <| NavTo link
                            , href link
                            ]
                            [ Html.text <| "#" ++ build.name ]
                        ]
    in
        [ Html.h3 [ class "man pas ansi-bright-black-bg" ] [ Html.text jobName ]
        , Html.ul [ class "builds-list" ]
            (case (Dict.get jobName buildDict) of
                Nothing ->
                    []

                -- never happens
                Just buildList ->
                    (List.map oneBuildToLi buildList)
            )
        ]


updateExpandedProperties : Model -> List (Cmd Msg)
updateExpandedProperties model =
    let
        filteredList =
            List.filter
                (isExpanded model.versionedUIStates)
                model.versionedResources.content
    in
        List.concatMap
            (fetchInputAndOutputs model)
            filteredList


isExpanded : Dict.Dict Int VersionUIState -> Concourse.VersionedResource -> Bool
isExpanded states versionedResource =
    let
        state =
            Dict.get versionedResource.id states
    in
        case state of
            Nothing ->
                False

            Just someState ->
                someState.expanded


fetchInputAndOutputs : Model -> Concourse.VersionedResource -> List (Cmd Msg)
fetchInputAndOutputs model versionedResource =
    let
        identifier =
            { teamName = model.resourceIdentifier.teamName
            , pipelineName = model.resourceIdentifier.pipelineName
            , resourceName = model.resourceIdentifier.resourceName
            , versionID = versionedResource.id
            }
    in
        [ fetchInputTo identifier
        , fetchOutputOf identifier
        ]


fetchResource : Concourse.ResourceIdentifier -> Cmd Msg
fetchResource resourceIdentifier =
    Task.attempt ResourceFetched <|
        Concourse.Resource.fetchResource resourceIdentifier


pauseResource : Concourse.ResourceIdentifier -> Concourse.CSRFToken -> Cmd Msg
pauseResource resourceIdentifier csrfToken =
    Task.attempt PausedToggled <|
        Concourse.Resource.pause resourceIdentifier csrfToken


unpauseResource : Concourse.ResourceIdentifier -> Concourse.CSRFToken -> Cmd Msg
unpauseResource resourceIdentifier csrfToken =
    Task.attempt PausedToggled <|
        Concourse.Resource.unpause resourceIdentifier csrfToken


fetchVersionedResources : Concourse.ResourceIdentifier -> Maybe Page -> Cmd Msg
fetchVersionedResources resourceIdentifier page =
    Task.attempt (VersionedResourcesFetched page) <|
        Concourse.Resource.fetchVersionedResources resourceIdentifier page


enableVersionedResource : Concourse.VersionedResourceIdentifier -> Concourse.CSRFToken -> Cmd Msg
enableVersionedResource versionedResourceIdentifier csrfToken =
    Task.attempt (VersionedResourceToggled versionedResourceIdentifier.versionID) <|
        Concourse.Resource.enableVersionedResource versionedResourceIdentifier csrfToken


disableVersionedResource : Concourse.VersionedResourceIdentifier -> Concourse.CSRFToken -> Cmd Msg
disableVersionedResource versionedResourceIdentifier csrfToken =
    Task.attempt (VersionedResourceToggled versionedResourceIdentifier.versionID) <|
        Concourse.Resource.disableVersionedResource versionedResourceIdentifier csrfToken


fetchInputTo : Concourse.VersionedResourceIdentifier -> Cmd Msg
fetchInputTo versionedResourceIdentifier =
    Task.attempt (InputToFetched versionedResourceIdentifier.versionID) <|
        Concourse.Resource.fetchInputTo versionedResourceIdentifier


fetchOutputOf : Concourse.VersionedResourceIdentifier -> Cmd Msg
fetchOutputOf versionedResourceIdentifier =
    Task.attempt (OutputOfFetched versionedResourceIdentifier.versionID) <|
        Concourse.Resource.fetchOutputOf versionedResourceIdentifier


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ Time.every (5 * Time.second) AutoupdateTimerTicked
        , Time.every Time.second ClockTick
        ]
