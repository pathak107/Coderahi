import { useContext } from "react";
import { PostContext } from "../../context/postContext";
import PostForm from "../Course/PostForm";

const PostModal = ({children}) => {
    const {state, actions}= useContext(PostContext)
    return (
        <div className={`modal ${state.isOpen?"modal-open":""}`} >
            <div className="modal-box relative">
                <a className="btn btn-sm btn-circle absolute right-2 top-2"
                    onClick={()=>{
                        actions.closeModal()
                    }}
                >
                    ✕
                </a>
                <PostForm/>
            </div>
        </div>
    );
}

export default PostModal;