import CourseCard from "../../components/courses/CourseCard";
import { getAllCourses } from "../../services/api";

const Courses = ({courses}) => {
    return (
        <>
            {courses.map((course)=>{
                return <CourseCard course={course}/>
            })}
            
        </>
    );
}

export async function getStaticProps(context) {
    const courses= await getAllCourses()
    return {
      props: {
        courses
      }, // will be passed to the page component as props
    }
  }
 
export default Courses;