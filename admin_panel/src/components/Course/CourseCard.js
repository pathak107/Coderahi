import CreateCourseForm from "./CreateCourseForm";

const CourseCard = () => {
    return (
        <>
            <input type="checkbox" id="my-modal-3" className="modal-toggle" />
            <div className="modal">
                <div className="modal-box relative">
                    <label htmlFor="my-modal-3" className="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
                    <h3 className="text-lg font-bold">Enter following information for the new course</h3>
                    <CreateCourseForm />
                    <div className="modal-action">
                        <label htmlFor="my-modal" className="btn">Create</label>
                    </div>
                </div>
            </div>

            <div className="card w-96 bg-base-200 shadow-xl mx-2 my-2">
                <div className="card-body">
                    <h2 className="card-title">Course</h2>
                    <p>Create a brand new course here or look up and edit any existing courses.</p>
                    <div className="card-actions justify-end">
                        <button className="btn btn-sm btn-primary">Show All</button>
                        <label htmlFor="my-modal-3" className="btn btn-sm btn-primary modal-button">New</label>
                    </div>
                </div>
            </div>
        </>

    );
}

export default CourseCard;