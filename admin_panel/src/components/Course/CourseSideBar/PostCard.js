import { FaPlus, FaTrashAlt, FaEdit } from 'react-icons/fa'
import { ConfirmModalCtx } from '../../../context/confirmModalCtx';
import { PostContext } from '../../../context/postContext';
import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';

const PostCard = ({course_id, post, provided, innerRef}) => {
    const postModalCtx = useContext(PostContext)
    const confirmModalCtx = useContext(ConfirmModalCtx)
    const navigate = useNavigate()

    return (
        <li {...provided.draggableProps} {...provided.dragHandleProps} ref={innerRef}>
            <div className='flex flex-row'>
                <p>{post.Order+1}.</p>
                <a className='hover:underline' onClick={() => {
                    navigate(`/course/${course_id}/post/${post.ID}`)
                }}>
                    {post.Title}
                </a>
                <a className='btn btn-xs'
                    onClick={() => {
                        confirmModalCtx.actions.openModal()
                        confirmModalCtx.actions.setOnYesAction({
                            action: () => postModalCtx.actions.deletePostAct(post.ID)
                        })

                    }}
                >
                    <FaTrashAlt />
                </a>
                <a className='btn btn-xs'
                    onClick={() => {
                        postModalCtx.actions.setEdit(true)
                        postModalCtx.actions.setPostID(post.ID)
                        postModalCtx.actions.setTitle(post.Title)
                        postModalCtx.actions.setDesc(post.Description)
                        postModalCtx.actions.openModal()
                    }}
                >
                    <FaEdit />
                </a>
            </div>
        </li>
    );
}

export default PostCard;