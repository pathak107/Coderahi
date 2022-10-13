
const PostForm = (props) => {
    return (
        <div className="">
            <form className="flex flex-col justify-center">
                <input type="text" placeholder="Title" className="input input-bordered w-full max-w-xs" />
                <textarea className="textarea textarea-bordered" placeholder="Description"></textarea>
            </form>
            <a href="#" className="btn">Save</a>
        </div>
    );
}

export default PostForm;