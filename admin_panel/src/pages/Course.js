import CourseForm from "../components/Course/CourseForm";
import CourseSidebar from "../components/Course/CourseSidebar";
import MainContent from "../components/MainContent";

function Course() {
    return (
        <div className="flex flex-row">
            <CourseSidebar/>
            <MainContent>
                {/* is just course_id is there in params , show the edit course form otherwise if section and post_id is also present, show that */}
                <CourseForm/>
            </MainContent>
        </div>
    );
}

export default Course;