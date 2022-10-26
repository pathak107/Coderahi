import { useContext, useState } from "react";
import { createPost, createSection, editPost, editSection } from '../../services/api_service'
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { PostContext } from "../../context/postContext";

const PostForm = () => {
    const { state, actions } = useContext(PostContext)

    const queryClient = useQueryClient()

    const mutation = useMutation(state.edit ? editPost : createPost, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
            actions.setTitle("")
            actions.setDesc("")
        },
    })

    return (
        <div className="">
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" value={state.title}
                    onChange={(e) => {
                        actions.setTitle(e.target.value)
                    }}
                />
                <textarea className="textarea textarea-bordered" placeholder="Description" value={state.desc}
                    onChange={(e) => {
                        actions.setDesc(e.target.value)
                    }}
                >
                </textarea>
                <span className="label-text">Publish</span> 
                <input type="checkbox" className="toggle" checked={state.published} onClick={()=>{actions.setPublished(!state.published)}}/>
            </form>
            <a className="btn"
                onClick={() => {
                    if (state.edit == true) {
                        mutation.mutate({
                            post_id: state.postID,
                            title: state.title,
                            desc: state.desc,
                            publish: state.published
                        }, "edit-post")
                    } else {
                        mutation.mutate({
                            title: state.title,
                            desc: state.desc,
                            section_id: state.sectionID
                        }, "create-post")
                    }
                    actions.closeModal()
                }
                }
            >
                Save
            </a>
        </div>
    );
}

export default PostForm;