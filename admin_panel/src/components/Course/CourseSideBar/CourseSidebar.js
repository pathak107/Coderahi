import { useContext} from 'react';
import { FaPlus } from 'react-icons/fa'
import { useNavigate, useParams } from 'react-router-dom';
import SectionModal from '../../Modals/SectionModal';
import { SectionContext } from '../../../context/sectionContext';
import ConfirmationModal from '../../Modals/ConfirmationModal';
import PostModal from '../../Modals/PostModal';
import  {DragDropContext, Droppable} from 'react-beautiful-dnd'
import SectionCard from './SectionCard';

function CourseSidebar({ course }) {
    const {course_id} = useParams()
    const navigate = useNavigate()
    const sectionModalCtx = useContext(SectionContext)


    return (
        <>
            <SectionModal/>
            <ConfirmationModal/>
            <PostModal/>

            <div className="h-screen overflow-auto overscroll-contain bg-base-900 w-1/5">
                <div className="content px-3">

                    {/* For Heading */}
                    <div className="course py-3">
                        <a className="text-2xl hover:underline cursor-pointer"
                            onClick={()=>{
                                navigate(`/course/${course_id}`)
                            }}
                        >
                            {course.Title}
                        </a>
                    </div>

                    <ul className="menu menu-horizontal bg-base-100 rounded-box place-items-center">
                        <li>
                            <a onClick={()=>{
                                sectionModalCtx.actions.setEdit(false)
                                sectionModalCtx.actions.setCourseID(course_id)
                                sectionModalCtx.actions.openModal()
                            }}>
                                <FaPlus />
                            </a>
                        </li>
                    </ul>

                    {course.Sections.map((section) => {
                        return  <SectionCard course_id={course_id} section={section} key={section.ID}/>
                    })}
                </div>
            </div>
        </>
    );

}

export default CourseSidebar