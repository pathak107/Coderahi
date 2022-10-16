const CourseCard = ({course}) => {
    return (
        <div className="card card-compact w-80 bg-base-300 shadow-xl">
            <figure><img src="https://placeimg.com/400/225/arch" alt="Shoes" /></figure>
            <div className="card-body">
                <div className="badge badge-secondary">Free</div>
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