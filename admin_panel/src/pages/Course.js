import CourseForm from "../components/Course/CourseForm";
import CourseSidebar from "../components/Course/CourseSideBar/CourseSidebar";
import MainContent from "../components/MainContent";
import { useParams } from 'react-router-dom';
import { getAllCategories, getCourseByIDWithSectionsAndPosts, getPostByID } from '../services/api_service';
import { useQuery, useQueryClient } from "@tanstack/react-query";
import SectionContextProvider from "../context/sectionContext";
import PostContextProvider from "../context/postContext";
import PostEditor from "../components/Post/PostEditor";
import EditorContextProvider from "../context/editorCtx";

function Course() {
    const { course_id, post_id } = useParams()

    const courseQuery = useQuery([`getOneCourse`], () => getCourseByIDWithSectionsAndPosts(course_id), {cacheTime: 0})
    const postQuery = useQuery([`getOnePost-${post_id}`], () => getPostByID(post_id), {enabled: post_id? true: false, cacheTime:0})
    const catQuery = useQuery([`getAllCategories`], () => getAllCategories())


    if (courseQuery.isLoading || catQuery.isLoading ||(post_id && postQuery.isLoading)) {
        return (
            <>
                <div className="radial-progress" />
            </>
        )
    }

    return (
        <div className="flex flex-row">
            <SectionContextProvider>
                <PostContextProvider>
                    <CourseSidebar course={courseQuery.data.data.data.course}/>
                </PostContextProvider>
            </SectionContextProvider>

            <MainContent>
                {/* is just course_id is there in params , show the edit course form otherwise if section and post_id is also present, show that */}
                {post_id? 
                    <EditorContextProvider 
                        initialEditorData={postQuery.data.data.data.post.MarkDown}
                    >
                        <PostEditor postID={post_id} />
                    </EditorContextProvider>
                    :
                    <EditorContextProvider initialEditorData={courseQuery.data.data.data.course.MarkDown}>
                        <CourseForm course={courseQuery.data.data.data.course} cats={catQuery.data.data.data.categories}/>
                    </EditorContextProvider> 
                }
                
            </MainContent>
        </div>
    );
}

export default Course;