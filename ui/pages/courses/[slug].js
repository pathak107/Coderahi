import { useRouter } from "next/router";
import CourseLessons from "../../components/courses/CourseLessons";
import PostBody from "../../components/posts/PostBody";
import { getAllCourses, getCourseBySlugWithSectionsAndPosts } from "../../services/api";

const CourseDetails = ({course}) => {
    const router =  useRouter()
    if (router.isFallback)  return <div>Loading...</div>;
    return (
        <>
            <div className="flex ">
                <CourseLessons course={course}/>
                <PostBody />
            </div>
        </>
    );
}

export async function getStaticPaths() {
    const courses = await getAllCourses()
    const courseSlugs=[]
    courses.forEach(course => {
        courseSlugs.push({ params: { slug: course.Slug } })
    });

    return {
      paths: courseSlugs,
      fallback: true,
    };
  }

export async function getStaticProps({params}) {
    const course= await getCourseBySlugWithSectionsAndPosts(params.slug)
    return {
      props: {
        course
      }, // will be passed to the page component as props
    }
  }
 
export default CourseDetails;