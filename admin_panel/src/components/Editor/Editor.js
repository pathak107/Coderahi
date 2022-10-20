import EditorJS from '@editorjs/editorjs';
import { useContext, useEffect, useRef, useState } from 'react';

import Header from '@editorjs/header';
import { EditorContext } from '../../context/editorCtx';

const EDITTOR_HOLDER_ID = 'editorjs';


function Editor({ initialData }) {
    const editorInstance = useRef(null)

    const { state, actions } = useContext(EditorContext)

    useEffect(() => {
        if (!editorInstance.current) {
            initEditor()
        }

        return () => {
            if (editorInstance.current) {
                editorInstance.current.destroy()
                editorInstance.current = null
            }
        }
    }, [])

    const initEditor = () => {
        const editor = new EditorJS({
            holder: EDITTOR_HOLDER_ID,
            logLevel: "ERROR",
            data: initialData,
            onReady: () => {
                editorInstance.current = editor
            },
            onChange: async (api, e) => {
                saveEditorData()
            },
            autofocus: true,
            tools: {
                header: Header
            }
        });
    }

    const saveEditorData = async () => {
        const res = await editorInstance.current.save()
        actions.setContent(res)
    }

    return (
        <>
            <div className='border-2 border-neutral-700 prose dark:prose-invert max-w-none'>
                <div id={EDITTOR_HOLDER_ID} />
            </div>
        </>
    );
}

export default Editor;
