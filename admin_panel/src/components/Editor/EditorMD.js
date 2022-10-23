import MDEditor from '@uiw/react-md-editor';
import { useContext, useEffect, useRef } from 'react';
import { EditorContext } from '../../context/editorCtx';
import {debounce} from 'lodash'

const EditorMD = ({post_id}) => {
    const {state, actions } = useContext(EditorContext)
    const debouncedAutoSave = useRef(
        debounce(async (md) => {
            console.log("Saving")
            actions.autoSaveFunc(post_id, md)
        }, 2000)
    ).current;

    useEffect(()=>{
        return ()=>{
            debouncedAutoSave.flush()
        }
    },[])

   
    return (
        <div className="w-full mr-2">
            <MDEditor
                value={state.markD}
                onChange={(val)=>{
                    actions.setMarkdown(val)
                    debouncedAutoSave(val)
                }}
                autoFocus={true}
                height={700}
            />
            {/* <MDEditor.Markdown source={state.content} style={{ whiteSpace: 'pre-wrap' }} /> */}
        </div>
    );
}

export default EditorMD;
