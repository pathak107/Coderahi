import CourseCard from "../components/Course/CourseCard";

function Home() {
    return (
        <div className="artboard-demo">
            <div className="flex flex-row px-2 py-2">
                <CourseCard/>
            <div className="card w-96 bg-base-200 shadow-xl mx-2 my-2">
                <div className="card-body">
                    <h2 className="card-title">Category</h2>
                    <p>Create a brand new category for your courses or look up and edit any existing categories.</p>
                    <div className="card-actions justify-end">
                        <button className="btn btn-sm btn-primary">Show All</button>
                        <button className="btn btn-sm btn-primary">New</button>
                    </div>
                </div>
            </div>
            </div>
        </div>
    );
}

export default Home;