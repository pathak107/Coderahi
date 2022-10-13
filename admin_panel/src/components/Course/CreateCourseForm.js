import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import {createCourse} from '../../services/api_service'
import {useNavigate} from 'react-router-dom'

const CreateCourseForm = () => {
    const [title, setTitle]=useState("")
    const [description, setDescription]=useState("")
    const [cost, setCost] = useState(0)

    const navigate = useNavigate()

    const mutation =useMutation(createCourse, {
        onError: (error, variables, context) => {
          // An error happened!
          console.log(error)
        },
        onSuccess: ({data}, variables, context) => {
          navigate(`/course/${data.data.course_id}`)
        },
      })

    return (
        <div className="">
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" value={title} className="input input-bordered w-full max-w-xs" onChange={(e)=>{
                    setTitle(e.target.value)
                }} />
                <input type="number" placeholder="Cost" className="input input-bordered w-full max-w-xs" value={cost}
                    onChange={(e)=>{
                        setCost(e.target.value)
                    }}
                 />
                <textarea className="textarea textarea-bordered" placeholder="Short Description" value={description}
                    onChange={(e)=>{
                        setDescription(e.target.value)
                    }}
                >
                </textarea>
            </form>
            <a href="#" onClick={()=>{
                mutation.mutate({
                    title,
                    desc: description,
                    cost
                }, "create-course")
            }} className="btn">Create</a>
        </div>
    );
}

export default CreateCourseForm;