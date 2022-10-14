import { createContext, useState } from "react";

export const EditorContext = createContext()

const EditorContextProvider = ({children, initialEditorData}) => {
    const initEditorData= !initialEditorData? "" : JSON.parse(initialEditorData)
    const [content, setContent] = useState(initEditorData)
    const providerVal = {
        state:{content},
        actions:{setContent},
    }
    return ( 
        <EditorContext.Provider value={providerVal}>
            {children}
        </EditorContext.Provider>
    );
}
 
export default EditorContextProvider;