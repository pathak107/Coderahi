import Editor from "../Editor/Editor";
import { useQuery } from "@tanstack/react-query";
import { getPostByID } from "../../services/api_service";
import EditorContextProvider, { EditorContext } from "../../context/editorCtx"
import { useContext } from "react";

const PostEditor = ({postID}) => {
    const { isLoading, isError, data, error } = useQuery(['getOnePost'], () => getPostByID(postID))
    const editorCtx = useContext(EditorContext)

    if (isLoading) {
        return (
            <>
                <div className="radial-progress" />
            </>
        )
    }

    // On success loading of data, editorCtx.setContent()

    return (
        <>
            <Editor/>
        </>
    );
}
 
export default PostEditor;