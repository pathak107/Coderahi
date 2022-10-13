import CourseForm from "../components/Course/CourseForm";
import CourseSidebar from "../components/Course/CourseSidebar";
import MainContent from "../components/MainContent";
import { useParams } from 'react-router-dom';
import { getCourseByIDWithSectionsAndPosts } from '../services/api_service';
import { useQuery } from "@tanstack/react-query";

function Course() {
    const {course_id} = useParams()
    const { isLoading, isError, data, error } = useQuery(['getOneCourse'],()=>getCourseByIDWithSectionsAndPosts(course_id))

    if (isLoading) {
        return (
            <>
                <div className="radial-progress" style={{"--value":70}}>70%</div>
            </>
        )
    }
    console.log(data)
    return (
        <div className="flex flex-row">
            <CourseSidebar course={data.data.data.course} isLoading={isLoading}/>
            <MainContent>
                {/* is just course_id is there in params , show the edit course form otherwise if section and post_id is also present, show that */}
                <CourseForm/>
            </MainContent>
        </div>
    );
}

export default Course;