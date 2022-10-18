import Link from "next/link";


const CourseDescription = ({ course }) => {

    return (
        <div className="w-4/6 h-screen overflow-y-auto overflow-x-hidden overscroll-contain p-4">
            <div className="px-10">
                <div className="pb-4 mb-6 border-b-2 border-gray-700">
                    <h1 className="text-6xl font-extrabold text-sky-400 mb-2">{course.Title}</h1>
                    <p className="text-xs">Instructor</p>
                    <h3 className="text-xl font-bold">Shubham Pathak</h3>
                    {course.Cost!=0 ? <div className="badge badge-primary">Free</div>
                    : <h2 className="text-4xl font-bold">₹ {course.Cost}</h2>
                    }
                </div>
                <article className="prose lg:prose-xl leading-normal" dangerouslySetInnerHTML={{ __html: course.DescHTML }} />
            </div>
        </div>
    );
}

export default CourseDescription;