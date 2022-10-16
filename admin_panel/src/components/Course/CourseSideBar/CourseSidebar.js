import { useContext } from 'react';
import { FaPlus } from 'react-icons/fa'
import { useNavigate, useParams } from 'react-router-dom';
import SectionModal from '../../Modals/SectionModal';
import { SectionContext } from '../../../context/sectionContext';
import ConfirmationModal from '../../Modals/ConfirmationModal';
import PostModal from '../../Modals/PostModal';
import { DragDropContext, Draggable, Droppable } from 'react-beautiful-dnd'
import SectionCard from './SectionCard';
import {useState} from 'react'
import {useQueryClient, useMutation} from '@tanstack/react-query'
import { updateOrderOfSection } from '../../../services/api_service';


function CourseSidebar({ course }) {
    const { course_id } = useParams()
    const navigate = useNavigate()
    const sectionModalCtx = useContext(SectionContext)
    const [sections, setSections] = useState(course.Sections)

    const queryClient = useQueryClient()
    const mutation = useMutation(updateOrderOfSection, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getOneCourse"])
        },
    })

    const changeOrderOfSections = (e) => {
        if (!e.destination) return
        const changeSectionID =  sections[e.source.index].ID
        const items = Array.from(sections);
        const [reorderedItem] = items.splice(e.source.index, 1);
        items.splice(e.destination.index, 0, reorderedItem);
        setSections(items)

        // make a request 
        mutation.mutate({section_id: changeSectionID, order: e.destination.index}, 'change-section-order')
    }


    return (
        <>
            <SectionModal />
            <ConfirmationModal />
            <PostModal />

            <div className="h-screen overflow-auto overscroll-contain bg-base-900 w-1/5">
                <div className="content px-3">

                    {/* For Heading */}
                    <div className="course py-3">
                        <a className="text-2xl hover:underline cursor-pointer"
                            onClick={() => {
                                navigate(`/course/${course_id}`)
                            }}
                        >
                            {course.Title}
                        </a>
                    </div>

                    <ul className="menu menu-horizontal bg-base-100 rounded-box place-items-center">
                        <li>
                            <a onClick={() => {
                                sectionModalCtx.actions.setEdit(false)
                                sectionModalCtx.actions.setCourseID(course_id)
                                sectionModalCtx.actions.openModal()
                            }}>
                                <FaPlus />
                            </a>
                        </li>
                    </ul>

                    <DragDropContext onDragEnd={changeOrderOfSections}>
                        <Droppable droppableId='section-droppable'>
                            {(provided) => (
                                <div {...provided.droppableProps} ref={provided.innerRef}>
                                    {sections.map((section, index) => {
                                        return  <Draggable key={section.ID} draggableId={section.ID.toString()} index={index}>
                                            {(provided)=>(
                                                <SectionCard course_id={course_id} section={section} innerRef={provided.innerRef} provided={provided} />
                                            )}   
                                        </Draggable>
                                    })}
                                    {provided.placeholder}
                                </div>
                            )}
                        </Droppable>
                    </DragDropContext>
                </div>
            </div>
        </>
    );

}

export default CourseSidebar