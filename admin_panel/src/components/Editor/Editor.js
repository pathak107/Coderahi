import EditorJS from '@editorjs/editorjs';
import { useEffect, useRef, useState } from 'react';

import Header from '@editorjs/header';

const DEFAULT_INITIAL_DATA = () => {
    return {
        "time": new Date().getTime(),
        "blocks": [
            {
                "type": "header",
                "data": {
                    "text": "This is my awesome editor!",
                    "level": 1
                }
            },
        ]
    }
}

const EDITTOR_HOLDER_ID = 'editorjs';


function Editor() {
    const editorInstance =  useRef(null)
    const [editorData, setEditorData] = useState(DEFAULT_INITIAL_DATA);

    useEffect(()=>{
        console.log("Asdsad")
        if (!editorInstance.current){
            console.log("As")
            initEditor()
        }
        return ()=>{
            editorInstance.current=null
        }
    },[])

    const initEditor= ()=>{
        const editor = new EditorJS({
            holder: EDITTOR_HOLDER_ID,
            logLevel: "ERROR",
            data: editorData,
            onReady:()=>{
                editorInstance.current= editor
            },
            onChange: async () => {
                const res= await saveEditorData()
                setEditorData(res)
            },
            autofocus: true,
            tools: {
                header: Header
            }
        });
    }

    const saveEditorData= async ()=>{
        const res= await editorInstance.current.save()
        console.log(res)
        return res
    }

    return (
        <>
            <div >
                <div id={EDITTOR_HOLDER_ID} ref={editorInstance}/>
            </div>

        </>
    );
}

export default Editor;
