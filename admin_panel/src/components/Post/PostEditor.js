import Editor from "../Editor/Editor";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { getPostByID, editPost, editPostBody } from "../../services/api_service";
import EditorContextProvider, { EditorContext } from "../../context/editorCtx"
import { useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const PostEditor = ({postID, initialData}) => {
    const queryClient = useQueryClient()
    const editorCtx = useContext(EditorContext)

    const mutation = useMutation(editPostBody, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: (data , variables, context) => {

        },
    })

    useEffect(()=>{
        console.log("post editor mounted")
        return ()=>{
            queryClient.invalidateQueries([`getOnePost-${postID}`])
        }
    },[])

    return (
        <>
            <div className="my-2 p-2 bg-gray-900">
                <button className="btn btn-sm"
                    onClick={()=>{
                        mutation.mutate({
                            body: editorCtx.state.content,
                            post_id: postID
                        }, "edit-post-body")
                    }}
                >
                    Save
                </button>
            </div>
            <Editor initialData={JSON.parse(initialData)}/>
        </>
    );
}
 
export default PostEditor;