import CourseForm from "../components/Course/CourseForm";
import CourseSidebar from "../components/Course/CourseSidebar";
import MainContent from "../components/MainContent";
import { useParams } from 'react-router-dom';
import { getCourseByIDWithSectionsAndPosts } from '../services/api_service';
import { useQuery } from "@tanstack/react-query";
import SectionContextProvider from "../context/sectionContext";
import PostContextProvider from "../context/postContext";
import PostEditor from "../components/Post/PostEditor";
import EditorContextProvider from "../context/editorCtx";

function Course() {
    const { course_id, post_id } = useParams()
    const { isLoading, isError, data, error } = useQuery(['getOneCourse'], () => getCourseByIDWithSectionsAndPosts(course_id))

    if (isLoading) {
        return (
            <>
                <div className="radial-progress" style={{ "--value": 70 }}>70%</div>
            </>
        )
    }
    console.log(data)
    return (
        <div className="flex flex-row">
            <SectionContextProvider>
                <PostContextProvider>
                    <CourseSidebar course={data.data.data.course} isLoading={isLoading} />
                </PostContextProvider>
            </SectionContextProvider>

            <MainContent>
                {/* is just course_id is there in params , show the edit course form otherwise if section and post_id is also present, show that */}
                {post_id? 
                    <EditorContextProvider>
                        <PostEditor postID={post_id}/>
                    </EditorContextProvider>
                    :
                    <EditorContextProvider initialEditorData={data.data.data.course.DescJson}>
                        <CourseForm course={data.data.data.course} />
                    </EditorContextProvider> 
                }
                
            </MainContent>
        </div>
    );
}

export default Course;