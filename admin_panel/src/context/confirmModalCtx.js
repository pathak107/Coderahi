import { createContext, useState } from "react";

export const ConfirmModalCtx = createContext()

const ConfirmModalCtxProvider = ({children}) => {
    const [isOpen, setIsOpen]=useState(false)
    const [onYesAction, setOnYesAction]=useState({
        action:()=>{}
    })
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
        state:{isOpen, onYesAction},
        actions:{toggleModal, openModal, closeModal, setOnYesAction},
    }
    return ( 
        <ConfirmModalCtx.Provider value={providerVal}>
            {children}
        </ConfirmModalCtx.Provider>
    );
}
 
export default ConfirmModalCtxProvider;