import { createContext, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {deleteSection} from "../services/api_service"

export const SectionContext = createContext()

const SectionContextProvider = ({children}) => {
    const [isOpen, setIsOpen]=useState(false)
    const [edit, setEdit]= useState(false)
    const [courseID, setCourseID]=useState(null)
    const [sectionID, setSectionID]=useState(null)
    const [title, setTitle]=useState("")
    const [desc, setDesc]=useState("")

    const queryClient = useQueryClient()
    const mutation = useMutation(deleteSection, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries([`getOneCourse`])
        },
    })

    const deleteSectionAct = (section_id)=>{
        mutation.mutate(section_id, "deleteSection")
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
        state:{isOpen, edit, courseID, sectionID, title, desc},
        actions:{toggleModal, openModal, closeModal, setEdit, setCourseID, setSectionID, setTitle, setDesc, deleteSectionAct},
    }
    return ( 
        <SectionContext.Provider value={providerVal}>
            {children}
        </SectionContext.Provider>
    );
}
 
export default SectionContextProvider;