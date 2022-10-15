import { createContext, useState } from "react";

export const EditorContext = createContext()

const EditorContextProvider = ({children, initialEditorData}) => {
    const [content, setContent] = useState(null)

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