import { useRouter } from "next/router";
import CourseLessons from "../../../../components/courses/CourseLessons";
import MainLayout from "../../../../components/layout/MainLayout";
import PostBody from "../../../../components/posts/PostBody";
import { getAllCourses, getCourseBySlugWithSectionsAndPosts, getPostBySlug } from "../../../../services/api";

const CourseDetails = ({ course, post }) => {
    const router = useRouter()
    
    if (router.isFallback) return <div>Loading...</div>;
    return (
        <>
            <MainLayout>
                <div className="flex my-4">
                    <CourseLessons course={course} isImageRendered={false} showBuyNow={false}/>
                    <PostBody post={post}/>
                </div>
            </MainLayout>

        </>
    );
}

export async function getStaticPaths() {
    const courses = await getAllCourses()
    const courseSlugs = []
    courses.forEach(course => {
        course.Sections.forEach(section=>{
            section.Posts.forEach(post=>{
                courseSlugs.push({ 
                    params: { 
                        course_slug: course.Slug,
                        post_slug: post.Slug
                    } 
                })
            })
        })
    });

    return {
        paths: courseSlugs,
        fallback: true,
    };
}

export async function getStaticProps({ params }) {
    const course = await getCourseBySlugWithSectionsAndPosts(params.course_slug)
    const post = await getPostBySlug(params.post_slug)
    return {
        props: {
            course,
            post
        }, // will be passed to the page component as props
    }
}

export default CourseDetails;