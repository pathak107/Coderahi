import EditorJS from '@editorjs/editorjs';
import { useContext, useEffect, useRef, useState } from 'react';
import editorjsCodeflask from '@calumk/editorjs-codeflask';
import CodeBox from '@bomdi/codebox';
import Quote from '@editorjs/quote';
import ImageTool from '@editorjs/image';

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
                header: Header,
                code: editorjsCodeflask,
                codeBox: {
                    class: CodeBox,
                    config: {
                        themeURL: 'https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@9.18.1/build/styles/dracula.min.css', // Optional
                        themeName: 'atom-one-dark', // Optional
                        useDefaultTheme: 'light' // Optional. This also determines the background color of the language select drop-down
                    }
                },
                quote: {
                    class: Quote,
                    inlineToolbar: true,
                    shortcut: 'CMD+SHIFT+O',
                    config: {
                        quotePlaceholder: 'Enter a quote',
                        captionPlaceholder: 'Quote\'s author',
                    },
                },
                image: {
                    class: ImageTool,
                    config: {
                        endpoints: {
                            byFile: 'http://localhost:8008/uploadFile', // Your backend file uploader endpoint
                            byUrl: 'http://localhost:8008/fetchUrl', // Your endpoint that provides uploading by Url
                        }
                    }
                }
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
