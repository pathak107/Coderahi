import { useContext, useState } from 'react';
import { FaPlus, FaTrashAlt, FaEdit } from 'react-icons/fa'
import { useNavigate, useParams } from 'react-router-dom';
import SectionModal from '../Modals/SectionModal';
import { SectionContext } from '../../context/sectionContext';
import { ConfirmModalCtx } from '../../context/confirmModalCtx';
import ConfirmationModal from '../Modals/ConfirmationModal';
import { PostContext } from '../../context/postContext';
import PostModal from '../Modals/PostModal';

function CourseSidebar({ course }) {
    const {course_id} = useParams()
    const navigate = useNavigate()
    const sectionModalCtx = useContext(SectionContext)
    const postModalCtx = useContext(PostContext)
    const confirmModalCtx = useContext(ConfirmModalCtx)


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
                        return  <div className="section" key={section.ID}>
                            <div className="section-title text-base font-medium flex flex-row">
                                <p>{section.Title}</p>
                                <a className="btn btn-sm self-end"
                                    onClick={()=>{
                                        postModalCtx.actions.setEdit(false)
                                        postModalCtx.actions.setSectionID(section.ID)
                                        postModalCtx.actions.openModal()
                                    }}
                                >
                                    <FaPlus />
                                </a>
                                <a className="btn btn-sm self-end"
                                    onClick={()=>{
                                        confirmModalCtx.actions.openModal()
                                        confirmModalCtx.actions.setOnYesAction({
                                            action: ()=>sectionModalCtx.actions.deleteSectionAct(section.ID)
                                        })
                                        
                                    }}
                                >
                                    <FaTrashAlt />
                                </a>
                                <a className="btn btn-sm self-end"
                                    onClick={()=>{
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
                                <ul className="menu bg-base-200 text-xs">
                                    {section.Posts.map((post) => {
                                        return <li key={post.ID}>
                                            <div className='flex flex-row' onClick={()=>{
                                                    navigate(`/course/${course_id}/post/${post.ID}`)
                                                }}>
                                                <a className='hover:underline' onClick={()=>{
                                                    navigate(`/course/${course_id}/post/${post.ID}`)
                                                }}>
                                                    {post.Title}
                                                </a>
                                                <a className='btn btn-xs'
                                                    onClick={()=>{
                                                        confirmModalCtx.actions.openModal()
                                                        confirmModalCtx.actions.setOnYesAction({
                                                            action: ()=>postModalCtx.actions.deletePostAct(post.ID)
                                                        })
                                                        
                                                    }}
                                                >
                                                    <FaTrashAlt />
                                                </a>
                                                <a className='btn btn-xs'
                                                    onClick={()=>{
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