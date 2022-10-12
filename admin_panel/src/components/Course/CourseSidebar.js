function CourseSidebar() {
    return (
        <>
            <div className="h-screen overflow-auto overscroll-contain bg-base-900 w-1/5">
                <div className="content px-3">

                    {/* For Heading */}
                    <div className="course py-3">
                        <a className="text-2xl">System Desing  from Zero to Hero</a>
                    </div>

                    <ul className="menu menu-horizontal bg-base-100 rounded-box place-items-center">
                        <li>
                            <a>+</a>
                        </li>
                        <li>
                            <a>
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                            </a>
                        </li>
                        <li>
                            <a>
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" /></svg>
                            </a>
                        </li>
                    </ul>

                    <div className="section">
                        <div className="section-title text-base font-medium flex flex-row content-between">
                            <p>Section Name</p>
                            <a className="btn btn-sm">+</a>
                        </div>
                        <div className="section-content">
                            <ul className="menu bg-base-200 text-xs">
                                <li><a>Post 1</a></li>
                                <li className="bordered"><a>Post 2</a></li>
                                <li><a>Post 3</a></li>
                            </ul>
                        </div>
                    </div>

                    <div className="section">
                        <div className="section-title text-base font-medium flex flex-row content-between">
                            <p>Section Name</p>
                            <a className="btn btn-sm">+</a>
                        </div>
                        <div className="section-content">
                            <ul className="menu bg-base-200 text-xs">
                                <li><a>Post 1</a></li>
                                <li className="bordered"><a>Post 2</a></li>
                                <li><a>Post 3</a></li>
                            </ul>
                        </div>
                    </div>

                    <div className="section">
                        <div className="section-title text-base font-medium flex flex-row content-between">
                            <p>Section Name</p>
                            <a className="btn btn-sm">+</a>
                        </div>
                        <div className="section-content">
                            <ul className="menu bg-base-200 text-xs">
                                <li><a>Post 1</a></li>
                                <li className="bordered"><a>Post 2</a></li>
                                <li><a>Post 3</a></li>
                            </ul>
                        </div>
                    </div>

                </div>
            </div>
        </>
    );

}

export default CourseSidebar