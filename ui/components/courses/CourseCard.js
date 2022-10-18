import Image from 'next/image'

const CourseCard = ({course}) => {
    return (
        <div className="card card-compact w-80 bg-base-300 shadow-xl">
            <Image
                    src={"http://localhost:8080/" + course.ImageURL}
                    alt="Picture of the Course"
                    height={720}
                    width={1280}
                />
            <div className="card-body">
                <div className="badge badge-primary">Free</div>
                <h2 className="card-title">
                    <a href={`/courses/${course.Slug}`}>{course.Title}</a>
                </h2>
                <p>{course.DescShort}</p>
                <div className="card-actions justify-end">
                </div>
            </div>
        </div>
    );
}

export default CourseCard;