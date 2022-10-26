import { useContext, useState, useEffect } from "react";
import Editor from "../Editor/Editor";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { editCourse } from "../../services/api_service";
import { EditorContext } from "../../context/editorCtx";
import ImageUpload from "../ImageUpload/ImageUpload";
import EditorMD from "../Editor/EditorMD";
import Creatable, { useCreatable } from 'react-select/creatable';

const CourseForm = ({course, cats}) => {
    const [title, setTitle]=useState(course.Title)
    const [desc, setDesc]=useState(course.DescShort)
    const [cost, setCost]=useState(course.Cost)
    const [published, setPublished]=useState(course.Published)
    

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

    const addCategory = (vals)=>{
        console.log(vals)
        setCategories(vals)
    }

    const getAllOptions = ()=>{
        const options=[]
        cats.forEach(cat => {
            options.push({ value: cat.ID.toString(), label: cat.Name })
        });
        return options
    }

    const getCourseSelectedOptions = () => {
        const selectedOptions=[]
        course.Categories.forEach(cat=>{
            selectedOptions.push({ value: cat.ID.toString(), label: cat.Name })
        })
        return selectedOptions
    }

    const [categories, setCategories]= useState(getCourseSelectedOptions())

    useEffect(() => {
        return () => {
            queryClient.invalidateQueries([`getOneCourse`])
        }
    }, [])

    return (
        <div className="">
            <h1 className="text-4xl">Course Details</h1>
            <ImageUpload imageURL={course.ImageURL} course_id={course.ID}/>
            <span className="label-text">Publish</span> 
            <input type="checkbox" className="toggle" checked={published} onClick={()=>{setPublished(!published)}}/>
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
                <Creatable
                    options={getAllOptions()} 
                    isMulti={true}
                    backspaceRemovesValue={true}
                    isSearchable={true}
                    isClearable={true}
                    onChange={addCategory}
                    tabSelectsValue={true}
                    value={categories}
                    theme='neutral170'
                />
            </form>
            <button className="btn"
                onClick={()=>{
                    mutation.mutate({
                        title,
                        desc,
                        course_id: course.ID,
                        markdown: editorCtx.state.markD,
                        html: editorCtx.state.html,
                        cost,
                        categories,
                        publish: published
                    }, "edit-course")
                }}
            >
                Save
            </button>
            <div className="w-full">
                <EditorMD />
            </div>
        </div>
    );
}

export default CourseForm;