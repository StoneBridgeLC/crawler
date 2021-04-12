package headline

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//curl --location --request GET 'https://apis.naver.com/commentBox/cbox/web_naver_list_jsonp.json?ticket=news&pool=cbox5&lang=ko&country=KR&objectId=news055%2C0000887008&pageSize=2&page=2' \
//--header 'authority: apis.naver.com' \
//--header 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
//--header 'sec-ch-ua-mobile: ?0' \
//--header 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36' \
//--header 'accept: */*' \
//--header 'sec-fetch-site: same-site' \
//--header 'sec-fetch-mode: no-cors' \
//--header 'sec-fetch-dest: script' \
//--header 'referer: https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=104&oid=015&aid=0004528764' \
//--header 'accept-language: ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7' \
//--header 'cookie: NNB=GYHJCKDIDPRV6; MM_NEW=1; NFS=2; MM_NOW_COACH=1; ASID=dc7494e1000001774d5cbaed00000054; _ga=GA1.2.757796585.1610107391; _ga_7VKFYR6RV1=GS1.1.1616388447.5.0.1616388447.60; nx_ssl=2; BMR=s=1618223855325&r=https%3A%2F%2Fm.blog.naver.com%2FPostView.nhn%3FblogId%3Dkwonsukmin%26logNo%3D221238775732%26proxyReferer%3Dhttps%3A%252F%252Fwww.google.com%252F&r2=https%3A%2F%2Fwww.google.com%2F'

type CommentResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Lang    string `json:"lang"`
	Country string `json:"country"`
	Result  struct {
		CommentList []struct {
			Ticket            string      `json:"ticket"`
			ObjectID          string      `json:"objectId"`
			CategoryID        string      `json:"categoryId"`
			TemplateID        string      `json:"templateId"`
			CommentNo         int64       `json:"commentNo"`
			ParentCommentNo   int64       `json:"parentCommentNo"`
			ReplyLevel        int         `json:"replyLevel"`
			ReplyCount        int         `json:"replyCount"`
			ReplyAllCount     int         `json:"replyAllCount"`
			ReplyPreviewNo    interface{} `json:"replyPreviewNo"`
			ReplyList         interface{} `json:"replyList"`
			ImageCount        int         `json:"imageCount"`
			ImageList         interface{} `json:"imageList"`
			ImagePathList     interface{} `json:"imagePathList"`
			ImageWidthList    interface{} `json:"imageWidthList"`
			ImageHeightList   interface{} `json:"imageHeightList"`
			CommentType       string      `json:"commentType"`
			StickerID         interface{} `json:"stickerId"`
			Sticker           interface{} `json:"sticker"`
			SortValue         int64       `json:"sortValue"`
			Contents          string      `json:"contents"`
			UserIDNo          string      `json:"userIdNo"`
			ExposedUserIP     interface{} `json:"exposedUserIp"`
			Lang              string      `json:"lang"`
			Country           string      `json:"country"`
			IDType            string      `json:"idType"`
			IDProvider        string      `json:"idProvider"`
			UserName          string      `json:"userName"`
			UserProfileImage  string      `json:"userProfileImage"`
			ProfileType       string      `json:"profileType"`
			ModTime           string      `json:"modTime"`
			ModTimeGmt        string      `json:"modTimeGmt"`
			RegTime           string      `json:"regTime"`
			RegTimeGmt        string      `json:"regTimeGmt"`
			SympathyCount     int         `json:"sympathyCount"`
			AntipathyCount    int         `json:"antipathyCount"`
			HideReplyButton   bool        `json:"hideReplyButton"`
			Status            int         `json:"status"`
			Mine              bool        `json:"mine"`
			Best              bool        `json:"best"`
			Mentions          interface{} `json:"mentions"`
			ToUser            interface{} `json:"toUser"`
			UserStatus        int         `json:"userStatus"`
			CategoryImage     interface{} `json:"categoryImage"`
			Open              bool        `json:"open"`
			LevelCode         interface{} `json:"levelCode"`
			Grades            interface{} `json:"grades"`
			Sympathy          bool        `json:"sympathy"`
			Antipathy         bool        `json:"antipathy"`
			SnsList           interface{} `json:"snsList"`
			MetaInfo          interface{} `json:"metaInfo"`
			Extension         interface{} `json:"extension"`
			AudioInfoList     interface{} `json:"audioInfoList"`
			Translation       interface{} `json:"translation"`
			Report            interface{} `json:"report"`
			MiddleBlindReport bool        `json:"middleBlindReport"`
			SpamInfo          interface{} `json:"spamInfo"`
			UserHomepageURL   interface{} `json:"userHomepageUrl"`
			Defamation        bool        `json:"defamation"`
			HiddenByCleanbot  bool        `json:"hiddenByCleanbot"`
			Score             float64     `json:"score"`
			ManagerLike       bool        `json:"managerLike"`
			Visible           bool        `json:"visible"`
			Manager           bool        `json:"manager"`
			Deleted           bool        `json:"deleted"`
			Expose            bool        `json:"expose"`
			ExposeByCountry   bool        `json:"exposeByCountry"`
			Virtual           bool        `json:"virtual"`
			Secret            bool        `json:"secret"`
			ProfileUserID     interface{} `json:"profileUserId"`
			BlindReport       bool        `json:"blindReport"`
			IDNo              string      `json:"idNo"`
			Blind             bool        `json:"blind"`
			ServiceID         interface{} `json:"serviceId"`
			UserBlocked       bool        `json:"userBlocked"`
			ContainText       bool        `json:"containText"`
			MaskedUserID      string      `json:"maskedUserId"`
			MaskedUserName    string      `json:"maskedUserName"`
			ValidateBanWords  bool        `json:"validateBanWords"`
			Anonymous         bool        `json:"anonymous"`
		} `json:"commentList"`
		PageModel struct {
			Page           int         `json:"page"`
			PageSize       int         `json:"pageSize"`
			IndexSize      int         `json:"indexSize"`
			StartRow       int         `json:"startRow"`
			EndRow         int         `json:"endRow"`
			TotalRows      int         `json:"totalRows"`
			StartIndex     int         `json:"startIndex"`
			TotalPages     int         `json:"totalPages"`
			FirstPage      int         `json:"firstPage"`
			PrevPage       int         `json:"prevPage"`
			NextPage       int         `json:"nextPage"`
			LastPage       int         `json:"lastPage"`
			Current        interface{} `json:"current"`
			Threshold      interface{} `json:"threshold"`
			MoveToLastPage bool        `json:"moveToLastPage"`
			MoveToComment  bool        `json:"moveToComment"`
			MoveToLastPrev bool        `json:"moveToLastPrev"`
		} `json:"pageModel"`
		ExposureConfig struct {
			Reason interface{} `json:"reason"`
			Status string      `json:"status"`
		} `json:"exposureConfig"`
		Count struct {
			Comment            int `json:"comment"`
			Reply              int `json:"reply"`
			ExposeCount        int `json:"exposeCount"`
			DelCommentByUser   int `json:"delCommentByUser"`
			DelCommentByMon    int `json:"delCommentByMon"`
			BlindCommentByUser int `json:"blindCommentByUser"`
			BlindReplyByUser   int `json:"blindReplyByUser"`
			Total              int `json:"total"`
		} `json:"count"`
		ListStatus string        `json:"listStatus"`
		Sort       string        `json:"sort"`
		BestList   []interface{} `json:"bestList"`
	} `json:"result"`
	Date string `json:"date"`
}

type Comment struct {
	Body string
	CreateTime time.Time
	UpdateTime time.Time
}

func SetLastCrawledTimeisNull() time.Time{
	return time.Unix(0, 0)
}

func scrapComments(client *http.Client, newsUrl string, lastCrawledTime time.Time) ([]Comment, error) {
	// news url에서 필요한 정보 추출
	u, err := url.Parse(newsUrl)
	if err != nil {
		return nil, err
	}

	// oid와 aid의 조합으로 댓글 API요청해야 함
	oid := u.Query().Get("oid")
	aid := u.Query().Get("aid")

	comments := make([]Comment, 0)
	// 모든 댓글을 확인하기 위해 "최신순" 정렬(sort=NEW) 후 1페이지부터 끝까지 요청
	curPage := 1
	for curPage != 0 {
		// build request
		// build request url
		reqUrl, err := url.Parse("https://apis.naver.com/commentBox/cbox/web_naver_list_jsonp.json")
		if err != nil {
			return nil, err
		}
		params := url.Values{}
		params.Add("ticket", "news")
		params.Add("lang", "ko")
		params.Add("country", "KR")
		params.Add("pageSize", "100")
		params.Add("page", strconv.Itoa(curPage))
		params.Add("pool", "cbox5")
		params.Add("objectId", fmt.Sprintf("news%s,%s",oid, aid))
		params.Add("sort", "NEW")
		reqUrl.RawQuery = params.Encode()

		// build request
		req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
		if err != nil {
			return nil, err
		}
		// set request header
		req.Header.Add("authority", "apis.naver.com")
		req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\"")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
		req.Header.Add("accept", "*/*")
		req.Header.Add("sec-fetch-site", "same-site")
		req.Header.Add("sec-fetch-mode", "no-cors")
		req.Header.Add("sec-fetch-dest", "script")
		req.Header.Add("referer", "https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=104&oid=015&aid=0004528764")
		req.Header.Add("accept-language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
		//req.Header.Add("cookie", "NNB=GYHJCKDIDPRV6; MM_NEW=1; NFS=2; MM_NOW_COACH=1; ASID=dc7494e1000001774d5cbaed00000054; _ga=GA1.2.757796585.1610107391; _ga_7VKFYR6RV1=GS1.1.1616388447.5.0.1616388447.60; nx_ssl=2; BMR=s=1618223855325&r=https%3A%2F%2Fm.blog.naver.com%2FPostView.nhn%3FblogId%3Dkwonsukmin%26logNo%3D221238775732%26proxyReferer%3Dhttps%3A%252F%252Fwww.google.com%252F&r2=https%3A%2F%2Fwww.google.com%2F")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("scrapComments : response statuscode is not 200\n%v", spew.Sdump(req))
		}

		// response body :
		//_callback({"success":true,"code":"1000","message":"요청을 성공적으로 처리하였습니다.","lang":"ko","country":"KR","result":{"commentList":[{"ticket":"news","objectId":"news055,0000887008","categoryId":"*","templateId":"default_politics","commentNo":2530331624,"parentCommentNo":2530331624,"replyLevel":1,"replyCount":0,"replyAllCount":0,"replyPreviewNo":null,"replyList":null,"imageCount":0,"imageList":null,"imagePathList":null,"imageWidthList":null,"imageHeightList":null,"commentType":"txt","stickerId":null,"sticker":null,"sortValue":1618226005643,"contents":"바보 총리 +  멍청이  대통령 =  나라 폭망 .  ..  국제정세 하나도 모르고 왕따 외교","userIdNo":"49ZY3","exposedUserIp":null,"lang":"ko","country":"KR","idType":"naver","idProvider":"naver","userName":"nabo****","userProfileImage":"","profileType":"naver","modTime":"2021-04-12T20:13:26+0900","modTimeGmt":"2021-04-12T11:13:26+0000","regTime":"2021-04-12T20:13:26+0900","regTimeGmt":"2021-04-12T11:13:26+0000","sympathyCount":13,"antipathyCount":0,"hideReplyButton":false,"status":0,"mine":false,"best":false,"mentions":null,"toUser":null,"userStatus":0,"categoryImage":null,"open":false,"levelCode":null,"grades":null,"sympathy":false,"antipathy":false,"snsList":null,"metaInfo":null,"extension":null,"audioInfoList":null,"translation":null,"report":null,"middleBlindReport":false,"spamInfo":null,"userHomepageUrl":null,"defamation":false,"hiddenByCleanbot":false,"score":0.0,"managerLike":false,"visible":true,"manager":false,"deleted":false,"expose":true,"exposeByCountry":false,"virtual":false,"secret":false,"profileUserId":null,"idNo":"49ZY3","blindReport":false,"blind":false,"serviceId":null,"userBlocked":false,"containText":true,"maskedUserId":"nabo****","maskedUserName":"na****","validateBanWords":false,"anonymous":false}],"pageModel":{"page":1,"pageSize":1,"indexSize":10,"startRow":1,"endRow":1,"totalRows":65,"startIndex":0,"totalPages":65,"firstPage":1,"prevPage":0,"nextPage":2,"lastPage":10,"current":null,"threshold":null,"moveToLastPage":false,"moveToComment":false,"moveToLastPrev":false},"exposureConfig":{"reason":null,"status":"COMMENT_ON"},"count":{"comment":65,"reply":3,"exposeCount":65,"delCommentByUser":4,"delCommentByMon":0,"blindCommentByUser":0,"blindReplyByUser":0,"total":68},"listStatus":"current","sort":"FAVORITE","bestList":[]},"date":"2021-04-12T13:55:58+0000"});
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyString := string(bodyBytes)
		bodyString = strings.ReplaceAll(bodyString, "_callback(", "")
		bodyString = strings.ReplaceAll(bodyString, ");", "")

		commentResponse := CommentResponse{}
		if err := json.Unmarshal([]byte(bodyString), &commentResponse); err != nil {
			return nil, err
		}

		// set next page
		curPage = commentResponse.Result.PageModel.NextPage
		for _, comm := range commentResponse.Result.CommentList {
			regTimeGmt, err := time.Parse(commentTimeLayout, comm.RegTimeGmt)
			if err != nil {
				return nil, err
			}

			// regTimeGmt보다 lastCrawledTime이 더 최근이면 이미 크롤링 된 댓글
			if lastCrawledTime.After(regTimeGmt) || lastCrawledTime.Equal(regTimeGmt) {
				continue
			}

			modTimeGmt, err := time.Parse(commentTimeLayout, comm.ModTimeGmt)
			if err != nil {
				return nil, err
			}

			c := Comment{}
			c.Body = comm.Contents
			c.CreateTime = regTimeGmt
			c.UpdateTime = modTimeGmt
			comments = append(comments, c)
		}
	}

	return comments, nil
}

const commentTimeLayout = "2006-01-02T15:04:05-0700"