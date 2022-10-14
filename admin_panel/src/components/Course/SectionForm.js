import { useContext, useState } from "react";
import { createSection, editSection } from '../../services/api_service'
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { SectionContext } from "../../context/sectionContext";

const SectionForm = () => {
    const { state, actions } = useContext(SectionContext)

    const queryClient = useQueryClient()

    const mutation = useMutation(state.edit ? editSection : createSection, {
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
            </form>
            <a className="btn"
                onClick={() => {
                    if (state.edit == true) {
                        mutation.mutate({
                            section_id: state.sectionID,
                            title: state.title,
                            desc: state.desc,
                        }, "edit-section")
                    } else {
                        mutation.mutate({
                            title: state.title,
                            desc: state.desc,
                            course_id: state.courseID
                        }, "create-section")
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

export default SectionForm;