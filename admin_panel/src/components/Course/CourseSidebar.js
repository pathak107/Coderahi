import { useState } from 'react';
import { FaPlus, FaTrashAlt, FaEdit } from 'react-icons/fa'
import SectionForm from './SectionForm';
import PostForm from './PostForm'
import { useParams } from 'react-router-dom';
import GeneralModal from '../Modals/GeneralModal';

function CourseSidebar({ course }) {
    const [modalState, setModalState]= useState(false)
    const {course_id} = useParams()
    const [sectionInfo, setSectionInfo] = useState({
        edit: false,
        sectionID: null,
        courseID: null,
        title: "",
        desc: ""
    })

    const [postInfo, setPostInfo] = useState({
        edit: false,
        title: "",
        description: "",
        sectionID: "",
        postID: ""
    })


    return (
        <>
            <GeneralModal open={modalState}>
                <SectionForm sectionInfo={sectionInfo} />
            </GeneralModal>

            <div className="h-screen overflow-auto overscroll-contain bg-base-900 w-1/5">
                <div className="content px-3">

                    {/* For Heading */}
                    <div className="course py-3">
                        <a className="text-2xl">{course.Title}</a>
                    </div>

                    <ul className="menu menu-horizontal bg-base-100 rounded-box place-items-center">
                        <li>
                            <a href='#section-modal' onClick={()=>{
                                setSectionInfo({
                                    edit:false,
                                    courseID:  course_id,
                                })
                            }}>
                                <FaPlus />
                            </a>
                        </li>
                    </ul>

                    {course.Sections.map((section) => {
                        return  <div className="section" key={section.ID}>
                            <div className="section-title text-base font-medium flex flex-row">
                                <p>{section.Title}</p>
                                <a href="#post-modal" className="btn btn-sm self-end"><FaPlus /></a>
                                <a href="#delete-confirmation-modal" className="btn btn-sm self-end"><FaTrashAlt /></a>
                                <a href="#section-modal" className="btn btn-sm self-end"
                                    onClick={()=>{
                                        setSectionInfo({
                                            edit:true,
                                            sectionID: section.ID,
                                            title: section.Title,
                                            desc: section.Description
                                        })
                                        setModalState(true)
                                    }}
                                >
                                    <FaEdit />
                                </a>
                            </div>
                            <div className="section-content">
                                <ul className="menu bg-base-200 text-xs">
                                    {section.Posts.map((post) => {
                                        return <li key={post.ID}>
                                            <div className='flex flex-row'>
                                                <a>{post.Title}</a>
                                                <a href="#delete-confirmation-modal" className='btn btn-xs'><FaTrashAlt /></a>
                                                <a href="#post-modal" className='btn btn-xs'
                                                    onClick={()=>{
                                            
                                                    }}
                                                >
                                                    <FaEdit />
                                                </a>
                                            </div>
                                        </li>
                                    })}
                                </ul>
                            </div>
                        </div>

                    })}
                </div>
            </div>
        </>
    );

}

export default CourseSidebar