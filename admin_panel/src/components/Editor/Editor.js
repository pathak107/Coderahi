import EditorJS from '@editorjs/editorjs';
import { useContext, useEffect, useRef, useState } from 'react';

import Header from '@editorjs/header';
import { EditorContext } from '../../context/editorCtx';

const EDITTOR_HOLDER_ID = 'editorjs';


function Editor() {
    const editorInstance =  useRef(null)

    const {state, actions}= useContext(EditorContext)

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
            data: state.content,
            onReady:()=>{
                editorInstance.current= editor
            },
            onChange: async () => {
                saveEditorData()
            },
            autofocus: true,
            tools: {
                header: Header
            }
        });
    }

    const saveEditorData= async ()=>{
        const res= await editorInstance.current.save()
        actions.setContent(res)
    }

    return (
        <>
            <div className='border-2 border-neutral-300 prose dark:prose-invert max-w-none'>
                <div id={EDITTOR_HOLDER_ID} ref={editorInstance}/>
            </div>

        </>
    );
}

export default Editor;
