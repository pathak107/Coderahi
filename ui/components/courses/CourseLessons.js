import Link from "next/link";
import { useRouter } from "next/router";
import Image from 'next/image'

const CourseLessons = ({ course, isImageRendered, showBuyNow }) => {
    const router = useRouter()
    const { post_slug } = router.query
    return (
        <div className="h-full p-4 w-3/12 bg-base-300">
            <div>
                {isImageRendered ? <Image
                    src={"http://localhost:8080/" + course.ImageURL}
                    alt="Picture of the Course"
                    height={720}
                    width={1280}
                /> : <></>}
                <h4 className="text-l">Lessons</h4>
                <Link  href={`/courses/${course.Slug}`}>
                    <h2 className="text-2xl font-extrabold text-sky-400 cursor-pointer">{course.Title}</h2>
                </Link>
                {showBuyNow ?
                    <Link href={`/courses/${course.Slug}/posts/${course.Sections[0].Posts[0].Slug}`}>
                        <button className="btn w-full my-2">Start</button>
                    </Link> :
                    <></>
                }

            </div>
            <div className="overflow-auto overscroll-contain bg-base-900">
                <div>
                    {course.Sections.map((section, index) => {
                        return <div className="p-2" key={section.ID}>
                            <div className="flex flex-row gap-x-3 items-center">
                                <p className="text-xs">{section.Order + 1}</p>
                                <h5 className="text-xl font-bold">{section.Title}</h5>
                            </div>
                            <div>
                                {section.Posts.map((post) => {
                                    return <div className="py-1 px-2 mx-3" key={post.ID}>
                                        <div className="flex flex-row">
                                            <div className={`w-0.5 mr-1 ${post.Slug === post_slug ? "bg-sky-400" : ""}`}></div>
                                            <Link href={`/courses/${course.Slug}/posts/${post.Slug}`}>
                                                <h3 className={`text-sm cursor-pointer hover:text-sky-400 ${post.Slug === post_slug ? "text-sky-400" : ""} `}>{post.Title}</h3>
                                            </Link>
                                        </div>
                                    </div>
                                })}
                            </div>
                        </div>

                    })}
                </div>
            </div>

        </div>
    );
}

export default CourseLessons;