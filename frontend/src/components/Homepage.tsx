
import { Link } from 'react-router-dom';

const HomePage: React.FC = () => {
    return (
        <> 
            <div className = "text-6xl mb-5">
                welcome to parcel
                <div className = "mb-5">
                    please login here:
                </div>
                <div className = "flex flex-col">
                    <Link to ="/signup">signup here</Link>
                </div>
            </div>
        </>
    );
}

export default HomePage;