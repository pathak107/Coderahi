import { useMutation } from "@tanstack/react-query";
import { createContext, useState } from "react";
import Showdown from "showdown";
import { editPostBody } from "../services/api_service";

export const EditorContext = createContext()

const EditorContextProvider = ({children, initialEditorData}) => {
    const [markD, setMarkD] = useState(initialEditorData? initialEditorData:"")
    const [html, setHTML]= useState("")
    const [autoSaving, setAutoSaving]= useState(false)
    const SD= new Showdown.Converter()

    const postMutation = useMutation(editPostBody, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
            setAutoSaving(false)
        },
        onSuccess: (data, variables, context) => {
            console.log(data)
            setAutoSaving(false)
        },
    })
    

    const setMarkdown= (md) =>{
        setMarkD(md)
        const newHTML= SD.makeHtml(md)
        setHTML(newHTML)
    }

    const autoSaveFunc = (post_id, md) =>{
        if (!post_id) return

        setAutoSaving(true)
        const newHTML= SD.makeHtml(md)
        setHTML(newHTML)
        postMutation.mutate({
            markdown: md,
            html: newHTML,
            post_id
        }, "edit-post-body")
    }
    

    const providerVal = {
        state:{markD, html, autoSaving},
        actions:{setMarkD, setMarkdown, setAutoSaving, autoSaveFunc},
    }
    return ( 
        <EditorContext.Provider value={providerVal}>
            {children}
        </EditorContext.Provider>
    );
}
 
export default EditorContextProvider;