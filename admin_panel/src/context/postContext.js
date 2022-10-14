import { createContext, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {deletePost} from "../services/api_service"

export const PostContext = createContext()

const PostContextProvider = ({children}) => {
    const [isOpen, setIsOpen]=useState(false)
    const [edit, setEdit]= useState(false)
    const [postID, setPostID]=useState(null)
    const [sectionID, setSectionID]=useState(null)
    const [title, setTitle]=useState("")
    const [desc, setDesc]=useState("")

    const queryClient = useQueryClient()
    const mutation = useMutation(deletePost, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
        },
    })

    const deletePostAct = (post_id)=>{
        mutation.mutate(post_id, "deletePost")
    }

    const toggleModal = ()=>{
        setIsOpen(!isOpen)
    }
    const openModal=()=>{
        setIsOpen(true)
    }
    const closeModal=()=>{
        setIsOpen(false)
    }

    const providerVal={
        state:{isOpen, edit, postID, sectionID, title, desc},
        actions:{toggleModal, openModal, closeModal, setEdit, setPostID, setSectionID, setTitle, setDesc, deletePostAct},
    }
    return ( 
        <PostContext.Provider value={providerVal}>
            {children}
        </PostContext.Provider>
    );
}
 
export default PostContextProvider;