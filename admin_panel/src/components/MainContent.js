import Editor from "./Editor/Editor";

function  MainContent(props) {
    return (
        <div className="w-4/5 h-screen overflow-y-auto overflow-x-hidden overscroll-contain">
            {props.children}
        </div>
    );
}

export default MainContent;
