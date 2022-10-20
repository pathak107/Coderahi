const Header = () => {
    return (
        <div className="navbar bg-base-300 border-b-2">
            <div className="flex-1">
                <a className="btn btn-ghost normal-case text-xl">Coderahi Admin</a>
            </div>
            <div className="flex-none">
                <ul className="menu menu-horizontal p-0">
                    <li><a>Item 1</a></li>
                    <li><a>Item 3</a></li>
                </ul>
            </div>
        </div>
    );
}

export default Header;