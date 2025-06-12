export default function Profile() {

    



   
    return (
        <div className="">
            <div className="">
                <h1 className="">sohachimi's Profile</h1>

                <button >
                    {/* {isPrivate ? "Set Public" : "Set Private"} */}
                </button>
            </div>

            <div className="">
                <h2 className="">User Information</h2>
                <ul className="">
                    <li>Email: </li>
                    <li>Full Name: </li>
                </ul>
            </div>

            {/* {canView ? ( */}
            <>
                <div className="">
                    <h2 className="">User Activity</h2>
                    <p>"No recent activity."</p>
                </div>

                <div className="">
                    <h2 className="">Posts</h2>
                    {/* {posts.length > 0 ? ( */}
                    <ul className="">
                        {/* {posts.map(post => ( */}
                        <li className="">
                            <p></p>
                            <span className=""></span>
                        </li>
                        {/* ))} */}
                    </ul>
                    {/* ) : ( */}
                    <p>No posts yet.</p>
                    {/* )} */}
                </div>

                <div className="">
                    <h2 className="">Followers</h2>
                    <ul className="">
                        {/* {followers.map(f => (
                <li key={f.id}>{f.username}</li>
              ))} */}
                    </ul>
                </div>

                <div className="">
                    <h2 className="">Following</h2>
                    <ul className="">
                        {/* {following.map(f => (
                <li key={f.id}>{f.username}</li>
              ))} */}
                    </ul>
                </div>
            </>
            {/* ) : ( */}
            <div className="">
                This profile is private. Only followers can view the details.
            </div>
            {/* )} */}
        </div>
    );
}
