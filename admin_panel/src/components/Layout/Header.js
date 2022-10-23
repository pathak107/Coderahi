import { useNavigate } from "react-router-dom";

const Header = () => {
    const navigate = useNavigate()
    return (
        <div className="navbar bg-base-300 border-b-2">
            <div className="flex-1">
                <a className="btn btn-ghost normal-case text-xl">Coderahi Admin</a>
            </div>
            <div className="flex-none">
                <ul className="menu menu-horizontal p-0">
                    <li><a onClick={()=>navigate("/")}>Home</a></li>
                    <li><a onClick={()=>navigate("/course")}>Courses</a></li>
                    <li><a onClick={()=>navigate("/categories")}>Categories</a></li>
                </ul>
            </div>
        </div>
    );
}

export default Header;