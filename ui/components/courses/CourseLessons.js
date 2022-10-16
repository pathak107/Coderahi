const CourseLessons = ({course}) => {
    return (
        <div className="h-screen overflow-auto overscroll-contain bg-base-900 w-1/5">
            <ul>
                {course.Sections.map((s)=>{
                    return <li>{s.Title}</li>
                })}
            </ul>
        </div>
    );
}
 
export default CourseLessons;