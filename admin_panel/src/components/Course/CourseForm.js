import Editor from "../Editor/Editor";

const CourseForm = () => {
    return (
        <div className="">
            <figure><img src="https://res.cloudinary.com/practicaldev/image/fetch/s--o_lXvfhK--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/i/a0gvmzph343m9wvjys6h.png"></img></figure>
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" />
                <input type="number" placeholder="Cost" className="input input-bordered w-full max-w-xs" />
                <textarea className="textarea textarea-bordered" placeholder="Short Description"></textarea>
            </form>
            <div className="w-screen">
                <Editor/>
            </div>
            <button className="btn">Save</button>
        </div>
    );
}

export default CourseForm;