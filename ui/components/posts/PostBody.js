const PostBody = ({post}) => {

    return (
        <div className="w-9/12 h-screen overflow-y-auto overflow-x-hidden overscroll-contain p-4">
            <article className="prose lg:prose-xl leading-normal" dangerouslySetInnerHTML={{ __html: post.HTMLBody}} />
        </div>
    );
}
 
export default PostBody;