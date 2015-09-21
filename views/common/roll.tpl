 
		<div >
			<div class="carousel slide" id="carousel-602369">
				<ol class="carousel-indicators">
					<li class="active" data-slide-to="0" data-target="#carousel-602369">
					</li>
					<li data-slide-to="1" data-target="#carousel-602369">
					</li>
					<li data-slide-to="2" data-target="#carousel-602369">
					</li>
				</ol>
				<div class="carousel-inner">
					<div class="item active">
						<img alt="Carousel Bootstrap First" src="{{.banner1_url}}"  style="width:100%;height:345px;"/>
						<div class="carousel-caption roll">
							<h4>
									<a href="/article/view/{{.banner1_art.Id}}.html"> {{.banner1_art.Title}} </a>
							</h4>
							<p>
							{{substr (html2str .banner1_art.Content) 0 200}}
							</p>
						</div>
					</div>
					<div class="item">
						 
						<img alt="Carousel Bootstrap First" src="{{.banner2_url}}"  style="width:100%;height:345px;"/>
			 
					
						<div class="carousel-caption roll">
							<h4>
								<a href="/article/view/{{.banner2_art.Id}}.html"> {{.banner2_art.Title}} </a>
							</h4>
							<p>
                           {{substr (html2str .banner2_art.Content) 0 200}}					
                    		</p>
						</div>
					</div>
					<div class="item">
						<img alt="Carousel Bootstrap First" src="{{.banner3_url}}"  style="width:100%;height:345px;"/>
						<div class="carousel-caption roll">
							<h4>
								<a href="/article/view/{{.banner3_art.Id}}.html"> {{.banner2_art.Title}} </a>
							</h4>
							<p>
                                 {{substr (html2str .banner3_art.Content) 0 200}}								
                           </p>
						</div>
					</div>
				</div> <a class="left carousel-control" href="#carousel-602369" data-slide="prev"><span class="glyphicon glyphicon-chevron-left"></span></a> <a class="right carousel-control" href="#carousel-602369" data-slide="next"><span class="glyphicon glyphicon-chevron-right"></span></a>
			</div>
		</div>
 