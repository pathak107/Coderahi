import { createContext, useState } from "react";

export const ModalContext = createContext()

const ModalContextProvider = ({children}) => {
    const [isOpen, setIsOpen]=useState(false)
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
        state:{isOpen},
        actions:{toggleModal, openModal, closeModal},
    }
    return ( 
        <ModalContext.Provider value={providerVal}>
            {children}
        </ModalContext.Provider>
    );
}
 
export default ModalContextProvider;