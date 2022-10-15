import { useContext, useState } from "react";
import Editor from "../Editor/Editor";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { editCourse } from "../../services/api_service";
import { EditorContext } from "../../context/editorCtx";
import ImageUpload from "../ImageUpload/ImageUpload";

const CourseForm = ({course, initialData}) => {
    const [title, setTitle]=useState(course.Title)
    const [desc, setDesc]=useState(course.DescShort)
    const [cost, setCost]=useState(course.Cost)

    const editorCtx = useContext(EditorContext)

    const queryClient = useQueryClient()

    const mutation = useMutation(editCourse, {
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
            <h1 className="text-4xl">Course Details</h1>
            <ImageUpload imageURL={course.ImageURL} course_id={course.ID}/>
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" value={title}
                    onChange={(e)=>{
                        setTitle(e.target.value)
                    }}
                />
                <input type="number" placeholder="Cost" className="input input-bordered w-full max-w-xs" value={cost}
                    onChange={(e)=>{
                        setCost(e.target.value)
                    }}
                />
                <textarea className="textarea textarea-bordered" placeholder="Short Description" value={desc}
                    onChange={(e)=>{
                        setDesc(e.target.value)
                    }}
                >
                </textarea>
            </form>
            <div className="w-screen">
                <Editor initialData={JSON.parse(initialData)}/>
            </div>
            <button className="btn"
                onClick={()=>{
                    mutation.mutate({
                        title,
                        desc,
                        course_id: course.ID,
                        body: editorCtx.state.content,
                        cost
                    }, "edit-course")
                }}
            >
                Save
            </button>
        </div>
    );
}

export default CourseForm;