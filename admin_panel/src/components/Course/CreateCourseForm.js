
const CreateCourseForm = () => {
    return (
        <div className="">
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" />
                <input type="number" placeholder="Cost" className="input input-bordered w-full max-w-xs" />
                <textarea className="textarea textarea-bordered" placeholder="Short Description"></textarea>
            </form>
        </div>
    );
}

export default CreateCourseForm;