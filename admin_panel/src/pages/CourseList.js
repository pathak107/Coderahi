import { getAllCourses } from '../services/api_service';
import { useQuery } from "@tanstack/react-query";
import CourseListCard from '../components/Course/CourseListCard';

const CourseList = () => {
    const { isLoading, isError, data, error } = useQuery(['getAllCourses'], getAllCourses)

    if (isLoading) {
        return (
            <>
                <div className="radial-progress" style={{ "--value": 70 }}>70%</div>
            </>
        )
    }

    return (
        <>
            <div className="grid grid-cols-4 gap-2">
                {data.data.data.courses.map((course) => {
                    return <CourseListCard course={course} key={course.ID}/>
                })}
            </div>
        </>
    );
}

export default CourseList;