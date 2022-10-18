import CourseCard from "../../components/courses/CourseCard";
import MainLayout from "../../components/layout/MainLayout";
import { getAllCourses } from "../../services/api";

const Courses = ({ courses }) => {
  return (
    <>
      <MainLayout>
        <div className="p-4 my-6">
          <div className="grid grid-cols-4 ">
            {courses.map((course) => {
              return <CourseCard course={course} key={course.ID}/>
            })}
          </div>
        </div>
      </MainLayout>
    </>
  );
}

export async function getStaticProps(context) {
  const courses = await getAllCourses()
  return {
    props: {
      courses
    }, // will be passed to the page component as props
  }
}

export default Courses;