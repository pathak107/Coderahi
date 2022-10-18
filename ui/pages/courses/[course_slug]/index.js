import { useRouter } from "next/router";
import CourseLessons from "../../../components/courses/CourseLessons";
import MainLayout from "../../../components/layout/MainLayout";
import { getAllCourses, getCourseBySlugWithSectionsAndPosts, getPostBySlug } from "../../../services/api";
import CourseDescription from "../../../components/courses/CourseDescription";

const CoursePurchase = ({ course, post }) => {
    const router = useRouter()
    
    if (router.isFallback) return <div>Loading...</div>;
    return (
        <>
            <MainLayout>
                <div className="flex my-4">
                    <CourseLessons course={course} isImageRendered={true} showBuyNow={true}/>
                    <CourseDescription course={course} />
                </div>
            </MainLayout>

        </>
    );
}

export async function getStaticPaths() {
    const courses = await getAllCourses()
    const courseSlugs = []
    courses.forEach(course => {
        courseSlugs.push({ 
            params: { 
                course_slug: course.Slug,
            } 
        })
    });

    return {
        paths: courseSlugs,
        fallback: true,
    };
}

export async function getStaticProps({ params }) {
    const course = await getCourseBySlugWithSectionsAndPosts(params.course_slug)
    return {
        props: {
            course,
        }, // will be passed to the page component as props
    }
}

export default CoursePurchase;