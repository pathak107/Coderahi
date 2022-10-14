import { getAllCourses } from '../services/api_service';
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from 'react-router-dom';

const CourseList = () => {
    const navigate = useNavigate()
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
                    return <div className="card w-80 bg-base-100 shadow-xl">
                        <figure><img src="https://placeimg.com/400/225/arch" alt="Shoes" /></figure>
                        <div className="card-body">
                                <h3 className="card-title cursor-pointer hover:underline"
                                    onClick={
                                        () => {
                                            navigate(`/course/${course.ID}`)
                                        }
                                    }
                                >
                                    {course.Title}
                                </h3>
                            <p className="text-sm">{course.DescShort}</p>
                            <div className="card-actions justify-end">
                                <div className="badge badge-outline">Backend</div>
                                <div className="badge badge-outline">Software Development</div>
                            </div>
                        </div>
                    </div>
                })}
            </div>
        </>
    );
}

export default CourseList;