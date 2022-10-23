import { useMutation, useQueryClient } from "@tanstack/react-query";
import { editPostBody } from "../../services/api_service";
import { EditorContext } from "../../context/editorCtx"
import { useContext, useEffect } from "react";
import EditorMD from "../Editor/EditorMD";

const PostEditor = ({ postID }) => {
    const queryClient = useQueryClient()
    const editorCtx = useContext(EditorContext)

    const mutation = useMutation(editPostBody, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
            editorCtx.actions.setAutoSaving(false)
        },
        onSuccess: (data, variables, context) => {
            console.log(data)
            editorCtx.actions.setAutoSaving(false)
        },
    })

    useEffect(() => {
        return () => {
            queryClient.invalidateQueries([`getOnePost-${postID}`])
        }
    }, [])

    return (
        <>
            <div className="my-2 p-2 bg-gray-900 flex flex-row justify-between items-center pr-6 pl-6">
                <button className="btn btn-sm"
                    onClick={() => {
                        mutation.mutate({
                            markdown: editorCtx.state.markD,
                            html: editorCtx.state.html,
                            post_id: postID
                        }, "edit-post-body")
                    }}
                >
                    Save
                </button>
                {
                    editorCtx.state.autoSaving ? 
                    <div className="flex flex-row item-center">
                        <p>Auto Saving...</p>
                        <progress className="progress progress-primary w-56" />
                    </div> : <></>
                }


            </div>
            <EditorMD post_id={postID}/>
        </>
    );
}

export default PostEditor;