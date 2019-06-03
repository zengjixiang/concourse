module SideBar.Pipeline exposing (pipeline)

import Concourse
import Message.Message exposing (DomID(..), Message(..))
import Routes
import SideBar.Styles as Styles
import SideBar.Views as Views


type alias PipelineScoped a =
    { a
        | teamName : String
        , pipelineName : String
    }


pipeline :
    { a
        | hovered : Maybe DomID
        , currentPipeline : Maybe (PipelineScoped b)
    }
    -> Concourse.Pipeline
    -> Views.Pipeline
pipeline session p =
    let
        isCurrent =
            case session.currentPipeline of
                Just cp ->
                    cp.pipelineName == p.name && cp.teamName == p.teamName

                Nothing ->
                    False

        pipelineId =
            { pipelineName = p.name, teamName = p.teamName }

        isHovered =
            session.hovered == Just (SideBarPipeline pipelineId)
    in
    { icon =
        if isCurrent || isHovered then
            Styles.Bright

        else
            Styles.Dim
    , link =
        { opacity =
            if isCurrent || isHovered then
                Styles.Bright

            else
                Styles.GreyedOut
        , rectangle =
            if isCurrent then
                Styles.Dark

            else if isHovered then
                Styles.Light

            else
                Styles.PipelineInvisible
        , href =
            Routes.toString <|
                Routes.Pipeline { id = pipelineId, groups = [] }
        , text = p.name
        , domID = SideBarPipeline pipelineId
        }
    }
