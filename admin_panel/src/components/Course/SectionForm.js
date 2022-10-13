import { useState } from "react";
import { createSection, editSection } from '../../services/api_service'
import { useMutation, useQueryClient } from "@tanstack/react-query";

const SectionForm = ({ sectionInfo }) => {
    console.log(sectionInfo)
    const [title, setTitle] = useState(sectionInfo.title)
    const [desc, setDesc] = useState(sectionInfo.desc)

    const queryClient =  useQueryClient()

    const mutation = useMutation(sectionInfo.edit? editSection: createSection, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
        },
    })

    return (
        <div className="">
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" value={title}
                    onChange={(e) => {
                        setTitle(e.target.value)
                    }}
                />
                <textarea className="textarea textarea-bordered" placeholder="Description" value={desc}
                    onChange={(e) => {
                        setDesc(e.target.value)
                    }}
                >
                </textarea>
            </form>
            <a href="#" className="btn"
                onClick={()=>{
                    if (sectionInfo.edit==true){
                        mutation.mutate({
                            section_id: sectionInfo.sectionID,
                            title,
                            desc,
                        }, "edit-section")
                    }else{
                        mutation.mutate({
                            title,
                            desc,
                            course_id: sectionInfo.courseID
                        }, "create-section")}
                    }
                }
            >
                Save
            </a>
        </div>
    );
}

export default SectionForm;