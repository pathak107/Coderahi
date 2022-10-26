import { useRef, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { uploadCourseImage } from "../../services/api_service";

const ImageUpload = ({imageURL, course_id}) => {
    const [url, setUrl]= useState("http://localhost:8080/"+imageURL)
    const fileInput= useRef()
    
    const queryClient = useQueryClient()
    const mutation = useMutation(uploadCourseImage, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
        },
    })

    const uploadImageToServer=(e)=>{
        const fileObj = e.target.files && e.target.files[0];
        if (!fileObj) {
          return;
        }
        e.target.value = null;
        setUrl(URL.createObjectURL(fileObj))

        //upload to server
        mutation.mutate({
            file: fileObj,
            course_id
        }, "upload-course-image")


    }

    return (
        <div className="card card-compact w-96 bg-base-100 shadow-xl" onClick={()=>{
            fileInput.current.click()
        }}>
            <figure><img src={url} alt="Course Image" /></figure>
            <div className="card-body">
                <div className="card-actions justify-end">
                <input type="file" name="file" accept="image/*" hidden ref={fileInput} onChange={uploadImageToServer}/>
                </div>
            </div>
        </div>
    );
}

export default ImageUpload;