import { FaPlus, FaTrashAlt, FaEdit } from 'react-icons/fa'
import { SectionContext } from '../../../context/sectionContext';
import { ConfirmModalCtx } from '../../../context/confirmModalCtx';
import { PostContext } from '../../../context/postContext';
import PostCard from './PostCard';
import { useContext, useState } from 'react';
import { DragDropContext, Draggable, Droppable } from 'react-beautiful-dnd';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateOrderOfPost } from '../../../services/api_service';

const SectionCard = ({ section, course_id, provided, innerRef }) => {
    const postModalCtx = useContext(PostContext)
    const confirmModalCtx = useContext(ConfirmModalCtx)
    const sectionModalCtx = useContext(SectionContext)

    const queryClient = useQueryClient()
    const mutation = useMutation(updateOrderOfPost, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
        },
    })

    const sortPosts = () => {
        section.Posts.sort((post1, post2) => {
            if (post1.Order < post2.Order) return -1;
            if (post1.Order > post2.Order) return 1;
            return 0;
        })
        return section.Posts
    }

    const [posts, setPosts] = useState(sortPosts())

    const changeOrderOfPosts = (e) => {
        if (!e.destination) return
        const changePostID =  posts[e.source.index].ID
        const items = Array.from(posts);
        const [reorderedItem] = items.splice(e.source.index, 1);
        items.splice(e.destination.index, 0, reorderedItem);
        setPosts(items)

        // make a request 
        mutation.mutate({post_id: changePostID, order: e.destination.index}, 'change-post-order')
    }

    return (
        <div className="section" {...provided.draggableProps} {...provided.dragHandleProps} ref={innerRef}>
            <div className="section-title text-base font-medium flex flex-row">
                <p>{section.Order+1} {section.Title}</p>
                <a className="btn btn-sm self-end"
                    onClick={() => {
                        postModalCtx.actions.setEdit(false)
                        postModalCtx.actions.setSectionID(section.ID)
                        postModalCtx.actions.openModal()
                    }}
                >
                    <FaPlus />
                </a>
                <a className="btn btn-sm self-end"
                    onClick={() => {
                        confirmModalCtx.actions.openModal()
                        confirmModalCtx.actions.setOnYesAction({
                            action: () => sectionModalCtx.actions.deleteSectionAct(section.ID)
                        })

                    }}
                >
                    <FaTrashAlt />
                </a>
                <a className="btn btn-sm self-end"
                    onClick={() => {
                        sectionModalCtx.actions.setEdit(true)
                        sectionModalCtx.actions.setSectionID(section.ID)
                        sectionModalCtx.actions.setTitle(section.Title)
                        sectionModalCtx.actions.setDesc(section.Description)
                        sectionModalCtx.actions.openModal()
                    }}
                >
                    <FaEdit />
                </a>
            </div>
            <div className="section-content">
                <DragDropContext onDragEnd={changeOrderOfPosts}>
                    <Droppable droppableId='post-cards'>
                        {(provided) => (
                            <ul className="menu bg-base-200 text-xs" {...provided.droppableProps} ref={provided.innerRef}>
                                {posts.map((post, index) => {
                                    return <Draggable key={post.ID} draggableId={post.ID.toString()} index={index}>
                                        {(provided) => (
                                            <PostCard course_id={course_id} post={post} innerRef={provided.innerRef} provided={provided} />
                                        )}
                                    </Draggable>

                                })}
                                {provided.placeholder}
                            </ul>
                        )}
                    </Droppable>
                </DragDropContext>

            </div>
        </div>
    );
}

export default SectionCard;